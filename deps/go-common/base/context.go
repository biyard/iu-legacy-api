package base

import (
	"context"
	"mime/multipart"

	"biyard.co/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Context struct {
	context.Context

	c         *gin.Context
	requestID string
	zlog      []logger.DebugLog
}

func NewBaseContextFromGin(c *gin.Context) *Context {
	reqID := c.GetHeader("x-request-id")
	if reqID == "" {
		reqID = uuid.New().String()
	}

	return &Context{
		Context:   c.Request.Context(),
		c:         c,
		requestID: reqID,
		zlog:      []logger.DebugLog{},
	}
}

func (r *Context) MultipartForm() (*multipart.Form, error) {
	return r.c.MultipartForm()
}

func (r *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	return r.c.SaveUploadedFile(file, dst)
}

func NewBaseContext() *Context {
	return &Context{
		Context:   context.Background(),
		requestID: uuid.New().String(),
		zlog:      []logger.DebugLog{},
	}
}

func (r *Context) XRequestID() string  { return r.requestID }
func (r *Context) XB3ParentID() string { return "" }
func (r *Context) XB3SpanID() string   { return "" }
func (r *Context) XB3TraceID() string  { return "" }
func (r *Context) XAccountID() string  { return "" }
func (r *Context) AddLog(stack string, msg []any) {
	// logger.Debugf(nil, "stack: %s, msg: %v\n", stack, msg)
	r.zlog = append([]logger.DebugLog{{
		Stack:   stack,
		Message: msg,
	}}, r.zlog...)
}

func (r *Context) Logs() []logger.DebugLog {
	return r.zlog
}
