package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
	"github.com/opentracing/opentracing-go"
	"github.com/sethvargo/go-retry"
	"github.com/valyala/fasthttp"
	"go.uber.org/ratelimit"
)

var (
	ErrNoMethod = errors.New("require method in request")
	ErrNoUrl    = errors.New("require url in request")
)

type RequestOperation struct {
	HTTPMethod string
	HTTPPath   string
	HTTPHeader http.Header
}

// Requester is a marker to extract *http.Request
type Requester interface {
	GetRequest() *http.Request
}

// RateLimiter is used to rate-limit some process, possibly across goroutines.
type RateLimiter interface {
	// Take waits for RPS
	Take() time.Time
}

// ClientOption sets parameters for the HttpClient
type ClientOption func(cli *HTTPClient)

// Cli sets HttpClient's http doer i.e fasthttp.Client
func Cli(client *fasthttp.Client) ClientOption {
	return func(cli *HTTPClient) {
		cli.cli = client
	}
}

// GlobalHooks sets HttpClient's global hooks
func GlobalHooks(hooks ...Hook) ClientOption {
	return func(cli *HTTPClient) {
		for _, h := range hooks {
			cli.Hooks.AddHook(h)
		}
	}
}

// GlobalPostHooks sets HttpClient's global hooks
func GlobalPostHooks(hooks ...PostHook) ClientOption {
	return func(cli *HTTPClient) {
		for _, h := range hooks {
			cli.Hooks.AddPostHook(h)
		}
	}
}

// ConstantBackoff sets HttpClient's constant backoff e.g) [10ms, 10ms, ...], where 10 is duration
func ConstantBackoff(maxRetries uint64, duration time.Duration) ClientOption {
	if duration <= 0 {
		panic(fmt.Errorf("duration must be greater than 0"))
	}
	return func(cli *HTTPClient) {
		cli.backoff = func() retry.Backoff {
			b := retry.NewConstant(duration)
			return retry.WithMaxRetries(maxRetries, b)
		}
	}
}

// ExponentialBackoff sets HttpClient's exponential backoff
// e.g) [10 * 1, 10 * 2, 10 * 4, 10 * 8 ...], where 10 is base duration
func ExponentialBackoff(maxRetries uint64, base time.Duration) ClientOption {
	if base <= 0 {
		panic(fmt.Errorf("base must be greater than 0"))
	}
	return func(cli *HTTPClient) {
		cli.backoff = func() retry.Backoff {
			b := retry.NewExponential(base)
			return retry.WithMaxRetries(maxRetries, b)
		}
	}
}

// FibonacciBackoff sets HttpClient's fibonacchi backoff
// e.g) [10 * 1, 10 * 1, 10 * 2, 10 * 3, ...], where 10 is base duration
func FibonacciBackoff(maxRetries uint64, base time.Duration) ClientOption {
	if base <= 0 {
		panic(fmt.Errorf("base must be greater than 0"))
	}
	return func(cli *HTTPClient) {
		cli.backoff = func() retry.Backoff {
			b := retry.NewFibonacci(base)

			return retry.WithMaxRetries(maxRetries, b)
		}
	}
}

// BackOff sets HttpClient's backoff such as with jitter. see retry.Backoff's docs
func BackOff(backoff retry.Backoff) ClientOption {
	return func(cli *HTTPClient) {
		cli.backoff = func() retry.Backoff {
			return backoff
		}
	}
}

// Limiter sets HttpClient's global ratelimiter not per host.
// if u use want to limit request per host, use limiter at hook's PreRequest
func Limiter(rate int) ClientOption {
	return func(cli *HTTPClient) {
		cli.limiter = ratelimit.New(rate)
	}
}

// Tracer sets opentracing.Tracer to propagate span if exist in context.
func Tracer(tracer opentracing.Tracer) ClientOption {
	return func(cli *HTTPClient) {
		cli.tracer = tracer
	}
}

// Endpoint sets HttpClient's endpoint
func Endpoint(endpoint string) ClientOption {
	return func(cli *HTTPClient) {
		cli.endpoint = endpoint
	}
}

// Headers sets HttpClient's global headers
func Headers(headers http.Header) ClientOption {
	return func(cli *HTTPClient) {
		cli.headers = headers
	}
}

func DefaultClient(endpoint string) *HTTPClient {
	cli, _ := NewHttpClient(ConstantBackoff(3, 1*time.Second),
		Limiter(1000), Headers(http.Header{
			"Content-Type": []string{"application/json"},
		}), Endpoint(endpoint))

	return cli
}

func SimpleHTTPRequest() ClientOption {
	return func(cli *HTTPClient) {
		cli.doRequest = cli.doSimpleRequest
	}
}

// NewHttpClient returns a new HttpClient with given options
func NewHttpClient(opts ...ClientOption) (*HTTPClient, error) {
	cli := HTTPClient{
		cli:     &fasthttp.Client{},
		limiter: ratelimit.NewUnlimited(),
		tracer:  opentracing.NoopTracer{},
		backoff: func() retry.Backoff {
			return &noopBackoff{}
		},
	}
	cli.doRequest = cli.doFastHttpRequest
	for _, opt := range opts {
		if opt != nil {
			opt(&cli)
		}
	}
	return &cli, nil
}

func (c *HTTPClient) WithOpts(opts ...ClientOption) *HTTPClient {
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
	return c
}

// noopBackoff used default httpclient's backoff i.e no retry
type noopBackoff struct{}

func (n *noopBackoff) Next() (next time.Duration, stop bool) {
	return 0, true
}

type HTTPClient struct {
	Hooks
	endpoint  string
	cli       *fasthttp.Client
	limiter   RateLimiter
	tracer    opentracing.Tracer
	backoff   func() retry.Backoff
	headers   http.Header
	doRequest func(ctx context.Context, opt *SendOptions) (int, []byte, error)
}

func (c *HTTPClient) doSimpleRequest(ctx context.Context, opt *SendOptions) (int, []byte, error) {
	logger.DebugH(ctx.(*base.Context), func() {
		if opt.Body != nil {
			logger.Debugf(ctx, "doSimpleRequest: %s %s", opt.Method, opt.Url, string(opt.Body))
		} else {
			logger.Debugf(ctx, "doSimpleRequest: %s %s", opt.Method, opt.Url)
		}
	})
	req, err := http.NewRequest(opt.Method, opt.Url, bytes.NewBuffer(opt.Body))
	if err != nil {
		return 0, nil, err
	}

	for k, v := range opt.Header {
		req.Header.Set(k, v[0])
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, respBody, nil
}

func (r *HTTPClient) Endpoint() string {
	return r.endpoint
}

func (c *HTTPClient) doFastHttpRequest(ctx context.Context, opt *SendOptions) (int, []byte, error) {
	logger.DebugH(ctx.(*base.Context), func() {
		if opt.Body != nil {
			logger.Debugf(ctx, "doFastHttpRequest: %s %s\n%v\n\n%s", opt.Method, opt.Url, opt.Header, string(opt.Body))
		} else {
			logger.Debugf(ctx, "doFastHttpRequest: %s %s\n%v", opt.Method, opt.Url, opt.Header)
		}
	})
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	}
	if err := opt.Validate(); err != nil {
		return 0, nil, err
	}
	if opt.Header == nil {
		opt.Header = make(http.Header)
	}
	var (
		statusCode    int
		respBodyBytes []byte
		attempt       = 0
		requireUnwrap = false
	)
	err := retry.Do(ctx, c.backoff(), func(ctx context.Context) error {
		// setup requests
		attempt++
		var (
			req  = fasthttp.AcquireRequest()
			resp = fasthttp.AcquireResponse()
		)
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
		opt.buildRequest(ctx, c.tracer, req)
		opt.Hooks.beforeRequest(ctx, req, attempt)

		c.limiter.Take() // wait from limiter

		// do request
		var err error
		deadline, ok := ctx.Deadline()

		logger.Debugf(ctx.(*base.Context), "before request: %+v", req)
		if ok {
			err = c.cli.DoDeadline(req, resp, deadline)
		} else {
			err = c.cli.Do(req, resp)
		}
		logger.Debugf(ctx.(*base.Context), "after request: %+v, %v", resp, err)

		if err != nil {
			requireUnwrap = true
			return retry.RetryableError(err) // mark as a retryable
		}

		opt.Hooks.afterRequest(ctx, resp)

		code := resp.StatusCode()
		statusCode = code

		logger.DebugH(ctx.(*base.Context), func() {
			logger.Debugf(ctx.(*base.Context), "status code: %d", code)
		})

		// success
		if opt.IsAcceptableCode(code) {
			// resp body must copied to use again stream in fasthttp.Response
			respBody := resp.Body()
			copied := make([]byte, len(respBody))
			copy(copied, respBody)
			respBodyBytes = copied
			logger.DebugH(ctx.(*base.Context), func() {
				logger.Debugf(ctx.(*base.Context), "req: %s", string(opt.Body))
				logger.Debugf(ctx.(*base.Context), "resp: %s", string(respBodyBytes))
			})
			return nil
		}
		// fail
		statusErr := NewStatusError(opt.Method, opt.Url, resp)
		if opt.IsRetryableCode(code) {
			requireUnwrap = true
			return retry.RetryableError(statusErr) // mark as a retryable
		}
		requireUnwrap = false
		return statusErr
	})
	if err != nil {
		if requireUnwrap {
			if unwrap := errors.Unwrap(err); unwrap != nil {
				err = unwrap
			}
		}
		return statusCode, []byte{}, err
	}
	return statusCode, respBodyBytes, nil
}

func (c *HTTPClient) NewDefaultSendOptions(method, url string, body []byte, header http.Header) *SendOptions {
	return &SendOptions{
		Method:          method,
		Url:             url,
		Body:            body,
		Header:          header,
		Hooks:           c.CloneHooks(),
		AcceptableCodes: NewDefaultAcceptableCodes(),
		RetryableCodes:  NewDefaultRetryableCodes(),
	}
}

func AppendQuery(url string, opt interface{}) string {
	if opt == nil {
		return url
	}
	v, err := query.Values(opt)
	if err != nil {
		return url
	}
	q := v.Encode()
	if q != "" {
		url += "?" + q
	}
	return url
}

func (r HTTPClient) makeRequest(method, uri string, input interface{}) (RequestOperation, interface{}) {
	ret := RequestOperation{
		HTTPMethod: method,
		HTTPPath:   uri,
	}

	if input == nil {
		return ret, nil
	}

	val := reflect.Indirect(reflect.ValueOf(input))
	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		return ret, input
	}
	// FIXME: it does not supports array type; just, supports struct type.

	n := typ.NumField()
	var res any
	body := map[string]any{}

	query := []string{}
	headers := http.Header{}

	for i := 0; i < n; i++ {
		f := typ.Field(i)
		v := val.Field(i)

		// NOTE: it will skip a field if untagged. Additionally, it treats form as query parameters only.
		if k, ok := f.Tag.Lookup("json"); ok {
			body[k] = v.Interface()
		} else if k, ok := f.Tag.Lookup("form"); ok {
			query = append(query, fmt.Sprintf("%v=%v", k, url.QueryEscape(fmt.Sprintf("%v", v.Interface()))))
		} else if k, ok := f.Tag.Lookup("uri"); ok {
			uri = strings.Replace(uri, ":"+k, url.PathEscape(fmt.Sprintf("%v", v.Interface())), -1)
		} else if k, ok := f.Tag.Lookup("header"); ok {
			headers[k] = []string{fmt.Sprintf("%v", v.Interface())}
		}
	}

	if len(query) > 0 {
		uri = strings.Join([]string{uri, strings.Join(query, "&")}, "?")
	}

	if len(body) > 0 {
		res = body
	}

	return RequestOperation{
		HTTPMethod: method,
		HTTPPath:   uri,
		HTTPHeader: headers,
	}, res
}

func (r HTTPClient) Patch(ctx context.Context, path string, input, output any) (int, error) {
	ro, body := r.makeRequest("PATCH", path, input)
	return r.do(ctx, ro, body, output)
}

func (r HTTPClient) Put(ctx context.Context, path string, input, output any) (int, error) {
	ro, body := r.makeRequest("PUT", path, input)
	return r.do(ctx, ro, body, output)
}

func (r HTTPClient) Post(ctx context.Context, path string, input, output any) (int, error) {
	ro, body := r.makeRequest("POST", path, input)
	return r.do(ctx, ro, body, output)
}

func (r HTTPClient) Get(ctx context.Context, path string, input, output any) (int, error) {
	ro, body := r.makeRequest("GET", path, input)
	return r.do(ctx, ro, body, output)
}

func (r HTTPClient) Delete(ctx context.Context, path string, input, output any) (int, error) {
	ro, body := r.makeRequest("DELETE", path, input)
	return r.do(ctx, ro, body, output)
}

func (r HTTPClient) do(ctx context.Context, opt RequestOperation, input, output any) (int, error) {
	if len(r.headers) > 0 {
		if opt.HTTPHeader == nil {
			opt.HTTPHeader = r.headers
		} else {
			for k, v := range r.headers {
				opt.HTTPHeader[k] = v
			}
		}
	}

	url := r.endpoint + opt.HTTPPath
	span := startSpan(ctx, 4)

	if span != nil {
		if opt.HTTPHeader == nil {
			opt.HTTPHeader = make(http.Header)
		}
		opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(opt.HTTPHeader))
	}

	var body []byte
	if input != nil {
		if v, ok := input.([]byte); ok && len(v) > 0 {
			body = v
		} else {
			b, err := json.Marshal(input)
			if err != nil {
				return 0, err
			}

			body = b
		}
	}
	sendOpts := r.NewDefaultSendOptions(opt.HTTPMethod, url, body, opt.HTTPHeader)
	code, res, err := r.doRequest(ctx, sendOpts)

	if err != nil || res == nil || len(res) == 0 || output == nil {
		return code, err
	}

	if output != nil {
		err = json.Unmarshal(res, &output)
		if err != nil {
			return code, err
		}
	}
	return code, nil
}

// startSpan returns a new opentracing.Span with caller's function name as operation name.
func startSpan(ctx context.Context, callerSkip int) opentracing.Span {
	if callerSkip < 2 {
		callerSkip = 2
	}
	pc := make([]uintptr, 15)
	n := runtime.Callers(callerSkip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	opname := frame.Function

	var request *http.Request

	switch v := ctx.(type) {
	case Requester:
		request = v.GetRequest()
		ctx = request.Context()
	case *gin.Context:
		request = v.Request
		ctx = request.Context()
	}

	// extract span from context.
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		return opentracing.StartSpan(opname, opentracing.ChildOf(parentSpan.Context()))
	}

	// if provide *gin.Context or *rest.RestContext, we extract span context from request header.
	if request != nil {
		carrier := opentracing.HTTPHeadersCarrier(request.Header)
		tracer := opentracing.GlobalTracer()

		if wireContext, err := tracer.Extract(opentracing.HTTPHeaders, carrier); err == nil {
			return opentracing.StartSpan(opname, opentracing.ChildOf(wireContext))
		}
	}
	return opentracing.StartSpan(opname)
}
