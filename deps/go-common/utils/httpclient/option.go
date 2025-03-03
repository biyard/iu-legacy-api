package httpclient

import (
	"context"
	"net/http"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/valyala/fasthttp"
)

const (
	defaultContentType = "application/json"
)

var (
	defaultAcceptableCodes = NewDefaultAcceptableCodes()
	defaultRetryableCodes  = NewDefaultRetryableCodes()
)

// NewDefaultRetryableCodes returns a new default retryable codes include
// [http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout]
func NewDefaultRetryableCodes() map[int]struct{} {
	return map[int]struct{}{
		http.StatusInternalServerError: {},
		http.StatusTooManyRequests:     {},
		http.StatusBadGateway:          {},
		http.StatusServiceUnavailable:  {},
		http.StatusGatewayTimeout:      {},
	}
}

// NewDefaultAcceptableCodes returns a new default acceptable 2xx codes
func NewDefaultAcceptableCodes() map[int]struct{} {
	return map[int]struct{}{
		http.StatusOK:                   {},
		http.StatusCreated:              {},
		http.StatusAccepted:             {},
		http.StatusNonAuthoritativeInfo: {},
		http.StatusNoContent:            {},
		http.StatusResetContent:         {},
		http.StatusPartialContent:       {},
		http.StatusMultiStatus:          {},
		http.StatusAlreadyReported:      {},
		http.StatusIMUsed:               {},
	}
}

// SendOptions represents options of request
type SendOptions struct {
	Method          string
	Url             string
	Body            []byte
	Header          http.Header
	Hooks           Hooks
	AcceptableCodes map[int]struct{}
	RetryableCodes  map[int]struct{}
}

// IsAcceptableCode returns a true if given status is acceptable, otherwise false.
// if empty AcceptableCodes then use defaultAcceptableCodes.
func (opt *SendOptions) IsAcceptableCode(code int) bool {
	acceptable := opt.AcceptableCodes
	if len(acceptable) == 0 {
		acceptable = defaultAcceptableCodes
	}
	_, ok := acceptable[code]
	return ok
}

// IsRetryableCode returns a true if given status is retryable, otherwise false.
// if empty RetryableCodes then use defaultRetryableCodes.
func (opt *SendOptions) IsRetryableCode(code int) bool {
	retryable := opt.RetryableCodes
	if len(retryable) == 0 {
		retryable = defaultRetryableCodes
	}
	_, ok := retryable[code]
	return ok
}

// Validate return a true if valid options otherwise false,
// where Method, Url are required.
func (opt *SendOptions) Validate() error {
	if opt.Method == "" {
		return ErrNoMethod
	}
	if opt.Url == "" {
		return ErrNoUrl
	}
	return nil
}

func (opt *SendOptions) buildRequest(ctx context.Context, tracer opentracing.Tracer, req *fasthttp.Request) {
	if req == nil || ctx == nil {
		return
	}
	req.Header.SetContentType(defaultContentType)
	if len(opt.Header) != 0 {
		for key, values := range opt.Header {
			switch key {
			// Content-Type must use Set to avoid multiple values
			case "Content-Type":
				req.Header.SetContentType(values[0])
			default:
				for _, v := range values {
					req.Header.Add(key, v)
				}
			}
		}
	}
	if !strings.HasPrefix(opt.Url, "http") {
		opt.Url = "http://" + opt.Url
	}
	req.SetRequestURI(opt.Url)
	req.Header.SetMethod(opt.Method)
	if len(opt.Body) != 0 {
		req.SetBody(opt.Body)
	}
	// if tracer != nil {
	// 	// span := opentracing.SpanFromContext(ctx)
	// 	if span != nil {
	// 		// ignore error
	// 		// tracer.Inject(span.Context(), opentracing.HTTPHeaders, &tracing.FastHTTPHeadersCarrier{Header: &req.Header})
	// 	}
	// }
}
