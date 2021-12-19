package controller

import (
	"github.com/l1f/blockornot/validator"
	"net/http"
)

func (c *Controllers) errorResponse(webCtx *WebContext, status int, message interface{}) {
	env := envelope{"error": message}

	err := c.writeJSON(webCtx.Response, status, env, nil)
	if err != nil {
		c.ctx.Logger.Println(webCtx.Request, err)
		webCtx.Response.WriteHeader(500)
	}
}

func (c *Controllers) RateLimitExceededResponse(webCtx *WebContext) {
	message := "rate limit exceeded"
	c.errorResponse(webCtx, http.StatusTooManyRequests, message)
}

func (c *Controllers) ServerErrorResponse(webCtx *WebContext, err error) {
	c.ctx.Logger.Println(webCtx.Request, err)
	message := "the server encountered a problem and could not process your Request"
	c.errorResponse(webCtx, http.StatusInternalServerError, message)
}

func (c *Controllers) NotFoundResponse(webCtx *WebContext) {
	message := "the requested resource could not be found"
	c.errorResponse(webCtx, http.StatusNotFound, message)
}

func (c *Controllers) MethodNotAllowedResponse(webCtx *WebContext) {
	message := "the %s method is not supported for this resource"
	c.errorResponse(webCtx, http.StatusMethodNotAllowed, message)
}

func (c *Controllers) badRequestResponse(webCtx *WebContext, err error) {
	c.errorResponse(webCtx, http.StatusBadRequest, err.Error())
}

func (c *Controllers) failedValidationResponse(webCtx *WebContext, err *validator.ValidationError) {
	c.errorResponse(webCtx, http.StatusUnprocessableEntity, err)
}

//nolint:golint,unused
func (c *Controllers) editConflictResponse(webCtx *WebContext) {
	message := "unable to update the record due to an edit conflict, please try again"
	c.errorResponse(webCtx, http.StatusConflict, message)
}

func (c *Controllers) InvalidCredentialsResponse(webCtx *WebContext) {
	message := "invalid authentication credentials"
	c.errorResponse(webCtx, http.StatusUnauthorized, message)
}

//nolint:golint,unused
func (c *Controllers) invalidAuthenticationTokenResponse(webCtx *WebContext) {
	webCtx.Response.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	c.errorResponse(webCtx, http.StatusUnauthorized, message)
}

func (c *Controllers) AuthenticationRequiredResponse(webCtx *WebContext) {
	message := "you must be authenticated to access this resource"
	c.errorResponse(webCtx, http.StatusUnauthorized, message)
}

func (c *Controllers) InactiveAccountResponse(webCtx *WebContext) {
	message := "your user account must be activated to use this resource"
	c.errorResponse(webCtx, http.StatusForbidden, message)
}

func (c *Controllers) NotPermittedResponse(webCtx *WebContext) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	c.errorResponse(webCtx, http.StatusForbidden, message)
}
