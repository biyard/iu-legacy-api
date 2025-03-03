package base

import (
	"net/http"
	"reflect"

	"biyard.co/common/logger"
	"github.com/gin-gonic/gin"
)

type RenderType int

const (
	Redirect RenderType = iota
)

type Renderer interface {
	RenderingType() RenderType
	Location() string
}

func makeGinHandler(handler any) gin.HandlerFunc {
	// Compatibility with gin.HandlerFunc
	if v, ok := handler.(gin.HandlerFunc); ok {
		return v
	}

	val := reflect.ValueOf(handler)
	if val.Kind() != reflect.Func {
		panic("handler must be a callable function")
	}

	typ := val.Type()

	argN := typ.NumIn()

	if argN < 1 {
		panic("handler must have a context as its first argument")
	} else if argN > 2 {
		panic("handler must have at most two arguments; (Context, RequestType)")
	}

	if typ.NumOut() != 2 {
		panic("handler must have two return values; (ResponseType, BaseError)")
	}
	retError := typ.Out(1)
	if retError.Kind() == reflect.Pointer {
		retError = retError.Elem()
	}

	if retError != reflect.TypeOf(Error{}) {
		panic("handler must have a BaseError as its second return value")
	}

	tags := make([]string, 0)
	tagMap := make(map[string]bool)
	arg := typ.In(1)
	isPointer := false
	if arg.Kind() == reflect.Pointer {
		arg = arg.Elem()
		isPointer = true
	}

	if kind := arg.Kind(); kind == reflect.Struct {
		for j := 0; j < arg.NumField(); j++ {
			for _, k := range []string{"json", "uri", "form", "header"} {
				tag := arg.Field(j).Tag.Get(k)

				if tag != "" && !tagMap[k] {
					tagMap[k] = true
					tags = append(tags, k)
				}
			}
		}
	} else if kind == reflect.Map {
		tags = append(tags, "json")
	}

	return func(c *gin.Context) {
		in := reflect.New(arg)
		v := in.Interface()

		for _, tag := range tags {
			err := error(nil)
			switch tag {
			case "json":
				err = c.ShouldBindJSON(v)
			case "uri":
				err = c.ShouldBindUri(v)
			case "form":
				err = c.ShouldBindQuery(v)
			case "header":
				err = c.ShouldBindHeader(v)
			}

			if err != nil {
				c.JSON(400, ErrRequestParsing.WithMessage(err.Error()))
				return
			}
		}

		ctx := NewBaseContextFromGin(c)
		if !isPointer {
			in = reflect.Indirect(in)
		}
		logger.Infos(ctx, in.Interface())

		outs := val.Call([]reflect.Value{reflect.ValueOf(ctx), in})

		if !outs[1].IsZero() {
			e := outs[1].Interface().(*Error)
			logger.Errorf(ctx, "Error: %+v", e.Error())
			mret := logger.ManipulateLogs(ctx, e)
			c.JSON(e.StatusCode, mret)
			return
		}
		ret := outs[0].Interface()

		if v, ok := ret.(Renderer); ok && v.RenderingType() == Redirect {
			c.Redirect(http.StatusMovedPermanently, v.Location())
			return
		}
		logger.Debugf(ctx, "Response: %+v", ret)
		mret := logger.ManipulateLogs(ctx, ret)
		c.JSON(200, mret)
	}
}
