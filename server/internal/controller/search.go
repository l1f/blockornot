package controller

import (
	"errors"

	"github.com/l1f/blockornot/internal/controller/dto"
)

func (c *Controllers) Search(ctx *WebContext) {
	token := ctx.Request.Header.Get("X-AUTH-TOKEN")
	secret := ctx.Request.Header.Get("X-AUTH-SECRET")

	query := ctx.Request.URL.Query().Get("query")
	if len(query) == 0 {
		c.badRequestResponse(ctx, errors.New("the query parameter \"query\" must be specified "))
		return
	}

	tweets, err := c.ctx.Logic.SearchTweets(dto.Access{Token: token, Secret: secret}, query, nil)
	if err != nil {
		c.ServerErrorResponse(ctx, err)
		return
	}

	if len(*tweets) == 0 {
		c.NotFoundResponse(ctx)
		return
	}

	err = c.writeJSON(ctx, 200, tweets, nil)
	if err != nil {
		c.ServerErrorResponse(ctx, err)
	}
}
