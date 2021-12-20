package middleware

import (
	"github.com/l1f/blockornot/internal/application"
	"github.com/l1f/blockornot/internal/controller"
	"net/http"
)

type Middleware struct {
	ctx         *application.Context
	controllers *controller.Controllers
}

func New(ctx *application.Context, controllers *controller.Controllers) Middleware {
	return Middleware{ctx: ctx, controllers: controllers}
}

func RouterWrapper(next controller.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		context := controller.WebContext{Response: writer, Request: request}
		next(&context)
	})
}

func RouterToControllerWrapper(next http.Handler) controller.Handler {
	return func(ctx *controller.WebContext) {
		next.ServeHTTP(ctx.Response, ctx.Request)
	}
}
