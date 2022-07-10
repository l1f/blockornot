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
	middlewares := middleware.New(ctx, &controllers)

	// not found response
	router.NotFound = controller.ContextWrapperNoParams(controllers.NotFoundResponse)
	// not allowed response
	router.MethodNotAllowed = controller.ContextWrapperNoParams(controllers.MethodNotAllowedResponse)

	// auth
	router.GET("/api/v1/auth", controller.ContextWrapper(controllers.GetTwitterOAuthUrl))
	router.POST("/api/v1/auth", controller.ContextWrapper(controllers.CompleteTwitterAuth))

	// search
	router.GET("/api/v1/search", controller.ContextWrapper(middlewares.MustAuthenticated(controllers.Search)))

	// users
	router.GET("/api/v1/users/:id", controller.ContextWrapper(
		middlewares.MustAuthenticated(controllers.GetUserByID)))
	router.DELETE("/api/v1/users/:id", controller.ContextWrapper(
		middlewares.MustAuthenticated(controllers.BlockUserByID)))

	return middleware.RouterWrapper(
		middlewares.RecoverPanic(
			middlewares.RequestLogger(
				middleware.RouterToControllerWrapper(router))))
}
