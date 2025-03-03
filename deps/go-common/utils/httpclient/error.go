package httpclient

import (
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
)

// StatusError is returned if success to request but received not acceptable status code
type StatusError struct {
	Method       string
	URL          string
	StatusCode   int
	Header       http.Header
	ResponseDump string
}

func (s *StatusError) Error() string {
	resp := s.ResponseDump
	if len(resp) > 100 {
		resp = resp[:100] + "..."
	}
	return fmt.Sprintf("StatusError{Method:%s, URL:%s, StatusCode:%d, Body:%s}", s.Method, s.URL, s.StatusCode, resp)
}

// NewStatusError returns a new StatusError from given method and url and resp
func NewStatusError(method, url string, resp *fasthttp.Response) *StatusError {
	statusErr := StatusError{
		Method:       method,
		URL:          url,
		StatusCode:   resp.StatusCode(),
		Header:       make(http.Header),
		ResponseDump: string(resp.Body()),
	}
	resp.Header.VisitAll(func(key, value []byte) {
		statusErr.Header.Add(string(key), string(value))
	})
	return &statusErr
}
