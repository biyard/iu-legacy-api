package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"
)

type Hook interface {
	// BeforeRequest is called before request
	BeforeRequest(ctx context.Context, req *fasthttp.Request, attempt int)
}

type PostHook interface {
	// AfterRequest is called after request
	AfterRequest(ctx context.Context, res *fasthttp.Response)
}

// Hooks represents collection of Hook
type Hooks struct {
	hooks     []Hook
	postHooks []PostHook
}

// AddHook adds a given hook and it does not contains any locks i.e not thread-safe.
// Because Hooks in HttpClient is read only after initialized,
// and Hooks in SendOptions created per request maybe thread-safe.
func (hs *Hooks) AddHook(h Hook) {
	if h != nil {
		hs.hooks = append(hs.hooks, h)
	}
}

func (hs *Hooks) AddPostHook(h PostHook) {
	if h != nil {
		hs.postHooks = append(hs.postHooks, h)
	}
}

func (hs Hooks) CloneHooks() Hooks {
	clone := Hooks{}
	clone.hooks = hs.hooks[:]
	clone.postHooks = hs.postHooks[:]
	return clone
}

func (hs Hooks) beforeRequest(ctx context.Context, req *fasthttp.Request, attempt int) {
	for _, h := range hs.hooks {
		h.BeforeRequest(ctx, req, attempt)
	}
}

func (hs Hooks) afterRequest(ctx context.Context, res *fasthttp.Response) {
	for _, h := range hs.postHooks {
		h.AfterRequest(ctx, res)
	}
}
