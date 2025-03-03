//go:build js
// +build js

package base

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"syscall/js"
)

func (r *Server) handleWebWorker(this js.Value, args []js.Value) any {
	handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) any {
		resolve := promiseArgs[0]
		// reject := promiseArgs[1]

		go func() {
			w := &webWorkerResponseWrite{}
			data := bytes.NewBuffer([]byte{})

			defer func() {
				if e := recover(); e != nil {
					err := fmt.Sprintf("Error on handleWebWorker(%v): %v", data, string(debug.Stack()))
					resolve.Invoke(err)
				}
			}()
			if len(args) < 2 {
				resolve.Invoke("must have at least 2 arguments")
			}

			method := strings.ToUpper(args[0].String())
			path := args[1].String()
			if len(args) > 2 {
				data = bytes.NewBufferString(args[2].String())
			}

			req, e := http.NewRequest(method, path, data)
			if e != nil {
				resolve.Invoke(e.Error())
			}

			r.g.ServeHTTP(w, req)

			resolve.Invoke(w.String())
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func (r *Server) StartWebWorker() error {
	js.Global().Set("goWork", js.FuncOf(r.handleWebWorker))
	<-make(chan bool)

	return nil
}

func startWebWorkerImpl(r *Server) error {
	js.Global().Set("goWork", js.FuncOf(r.handleWebWorker))
	<-make(chan bool)

	return nil
}

func init() {
	startWebWorker = startWebWorkerImpl
}
