package middleware

import (
	"github.com/l1f/blockornot/internal/controller"
)

func (m Middleware) RequestLogger(next controller.Handler) controller.Handler {
	return func(ctx *controller.WebContext) {
		m.ctx.Logger.Info.Printf("%s - %s", ctx.Request.Method, ctx.Request.URL)
		next(ctx)
	}
}
