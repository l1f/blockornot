package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/l1f/blockornot/internal/application"
	"github.com/l1f/blockornot/internal/controller"
	"github.com/l1f/blockornot/internal/middleware"
	"net/http"
)

func New(ctx *application.Context) http.Handler {
	router := httprouter.New()
	controllers := controller.New(ctx)
	_ = middleware.New(ctx, &controllers)

	// not found response
	router.NotFound = controller.ContextWrapperNoParams(controllers.NotFoundResponse)
	// not allowed response
	router.MethodNotAllowed = controller.ContextWrapperNoParams(controllers.MethodNotAllowedResponse)

	return router
}
