package middleware

import (
	"fmt"
	"github.com/l1f/blockornot/internal/controller"
)

func (m Middleware) RecoverPanic(next controller.Handler) controller.Handler {
	return func(ctx *controller.WebContext) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Response.Header().Set("Connection", "close")
				m.controllers.ServerErrorResponse(ctx, fmt.Errorf("%s", err))
			}
		}()
		next(ctx)
	}
}
