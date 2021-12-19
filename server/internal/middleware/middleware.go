package middleware

import (
	"github.com/l1f/blockornot/internal/application"
	"github.com/l1f/blockornot/internal/controller"
)

type Middleware struct {
	ctx         *application.Context
	controllers *controller.Controllers
}

func New(ctx *application.Context, controllers *controller.Controllers) Middleware {
	return Middleware{ctx: ctx, controllers: controllers}
}
