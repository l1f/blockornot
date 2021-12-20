package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/l1f/blockornot/internal/application"
	"github.com/l1f/blockornot/internal/controller"
	"github.com/l1f/blockornot/internal/middleware"
)

func router(ctx *application.Context) http.Handler {
	router := httprouter.New()
	controllers := controller.New(ctx)
	_ = middleware.New(ctx, &controllers)

	// not found response
	router.NotFound = controller.ContextWrapperNoParams(controllers.NotFoundResponse)
	// not allowed response
	router.MethodNotAllowed = controller.ContextWrapperNoParams(controllers.MethodNotAllowedResponse)

	router.GET("/api/v1/auth", controller.ContextWrapper(controllers.GetTwitterOAuthUrl))

	return router
}
