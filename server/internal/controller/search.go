package controller

import (
	"errors"

	"github.com/l1f/blockornot/internal/controller/dto"
)

func (c *Controllers) Search(ctx *WebContext) {
	token, secret, err := c.getAuthHeader(ctx)
	if err != nil {
		c.badRequestResponse(ctx, err)
		return
	}

	query := ctx.Request.URL.Query().Get("query")
	if len(query) == 0 {
		c.badRequestResponse(ctx, errors.New("the query parameter \"query\" must be specified "))
	}

	tweets, err := c.ctx.Logic.SearchTweets(dto.Access{Token: token, Secret: secret}, query, nil)
	if err != nil {
		c.ServerErrorResponse(ctx, err)
	}

	err = c.writeJSON(ctx, 200, tweets, nil)
}