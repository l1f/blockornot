package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type WebContext struct {
	Response   http.ResponseWriter
	Request    *http.Request
	Parameters httprouter.Params
}

type Handler func(ctx *WebContext)

func ContextWrapper(handler Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		handler(&WebContext{
			Response:   writer,
			Request:    request,
			Parameters: params,
		})
	}
}

func ContextWrapperNoParams(handler Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler(&WebContext{
			Response:   writer,
			Request:    request,
			Parameters: nil,
		})
	}
}
