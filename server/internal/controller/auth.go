package controller

import (
	"github.com/l1f/blockornot/internal/controller/dto"
)

func (c *Controllers) GetTwitterOAuthUrl(ctx *WebContext) {
	request, err := c.ctx.Logic.TwitterLoginInit()
	if err != nil {
		c.ServerErrorResponse(ctx, err)
		return
	}

	c.setAuthHeader(ctx, request.Token, request.Secret)

	err = c.writeJSON(ctx, 200, dto.Request{Url: request.Url}, nil)
}

func (c *Controllers) CompleteTwitterOAuthUrl(ctx *WebContext) {
	var pin dto.Pin
	err := c.readJSON(ctx, &pin)
	if err != nil {
		c.badRequestResponse(ctx, err)
		return
	}

	token, secret, err := c.getAuthHeader(ctx)
	if err != nil {
		c.badRequestResponse(ctx, err)
		return
	}

	access, err := c.ctx.Logic.TwitterLoginResolve(dto.Request{
		Token:  token,
		Secret: secret,
	}, pin.Pin)

	if err != nil {
		c.ServerErrorResponse(ctx, err)
		return
	}

	c.setAuthHeader(ctx, access.Token, access.Secret)

	err = c.writeJSON(ctx, 200, nil, nil)
	if err != nil {
		c.ServerErrorResponse(ctx, err)
		return
	}
}
