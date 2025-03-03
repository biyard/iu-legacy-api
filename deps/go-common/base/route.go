package base

import (
	"net/http"

	"biyard.co/common/logger"
	"github.com/gin-gonic/gin"
)

type Route struct {
	rg *gin.RouterGroup
}

func NewRoute(rg *gin.RouterGroup) *Route {
	return &Route{
		rg: rg,
	}
}

func (r *Route) Group(relativePath string, router Router) {
	router.OnInit(NewRoute(r.rg.Group(relativePath)))
}

func (r *Route) Any(relativePath string, handler any) {
	logger.Debugf(nil, "[API] Any %s", relativePath)
	if v, ok := handler.(gin.HandlerFunc); ok {
		logger.Debugf(nil, "[Gin Handler] Any %s%s", r.rg.BasePath(), relativePath)
		r.rg.Any(relativePath, v)
	}
	r.rg.Any(relativePath, makeGinHandler(handler))
}

func (r *Route) GET(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodGet, relativePath, handler)
}

func (r *Route) POST(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodPost, relativePath, handler)
}

func (r *Route) DELETE(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodDelete, relativePath, handler)
}

func (r *Route) PATCH(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodPatch, relativePath, handler)
}

func (r *Route) PUT(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodPut, relativePath, handler)
}

func (r *Route) OPTIONS(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodOptions, relativePath, handler)
}

func (r *Route) HEAD(relativePath string, handler any) {
	r.registerMethodHandler(http.MethodHead, relativePath, handler)
}

func (r *Route) registerMethodHandler(method, relativePath string, handler any) {
	logger.Debugf(nil, "[API] %s %s%s", method, r.rg.BasePath(), relativePath)
	if v, ok := handler.(gin.HandlerFunc); ok {
		logger.Debugf(nil, "[Gin Handler] %s %s%s", method, r.rg.BasePath(), relativePath)
		r.rg.Handle(method, relativePath, v)
		return
	}
	r.rg.Handle(method, relativePath, makeGinHandler(handler))
}
