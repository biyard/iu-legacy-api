package base

import (
	"fmt"
	"net/http"
)

// webWorkerResponseWrite ...
type webWorkerResponseWrite struct {
	statusCode int
	data       []byte
}

func (r *webWorkerResponseWrite) Header() http.Header {
	return http.Header{}
}

func (r *webWorkerResponseWrite) Write(body []byte) (int, error) {
	r.data = body

	return len(body), nil
}

func (r *webWorkerResponseWrite) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func (r *webWorkerResponseWrite) String() string {
	if r.statusCode >= 200 && r.statusCode < 300 {
		return fmt.Sprintf(`{"statusCode": %d, "result": %s}`, r.statusCode, string(r.data))
	}

	return fmt.Sprintf(`{"statusCode": %d, "error": %s}`, r.statusCode, string(r.data))
}

func (r *webWorkerResponseWrite) IsError() bool {
	return (r.statusCode < 200 || r.statusCode >= 300)
}
