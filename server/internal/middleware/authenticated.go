package middleware

import "github.com/l1f/blockornot/internal/controller"

func (m Middleware) MustAuthenticated(next controller.Handler) controller.Handler {
	return func(ctx *controller.WebContext) {
		token := ctx.Request.Header.Get("X-AUTH-TOKEN")
		secret := ctx.Request.Header.Get("X-AUTH-SECRET")

		if secret == "" || token == "" {
			m.controllers.InvalidCredentialsResponse(ctx)
			return
		}

		next(ctx)
	}
}
