package middleware

import (
	"github.com/l1f/blockornot/internal/controller"
)

func (m Middleware) RequestLogger(next controller.Handler) controller.Handler {
	return func(ctx *controller.WebContext) {
		m.ctx.Logger.Info().Msgf("%s - %s", ctx.Request.Method, ctx.Request.URL)
		next(ctx)
	}
}
