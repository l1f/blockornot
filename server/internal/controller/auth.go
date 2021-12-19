package controller

import (
	"net/http"
)

func (c *Controllers) GetTwitterOAuthUrl(ctx *WebContext) {
	request, err := c.ctx.Logic.TwitterLoginInit()
	if err != nil {
		c.errorResponse(ctx, http.StatusInternalServerError, err)
	}

	err = c.writeJSON(ctx, 200, &request, nil)
}
