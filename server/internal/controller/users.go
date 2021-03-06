package controller

import (
	"errors"
	"github.com/l1f/blockornot/internal/controller/dto"
	"strconv"
)

func (c *Controllers) GetUserByID(ctx *WebContext) {
	token := ctx.Request.Header.Get("X-AUTH-TOKEN")
	secret := ctx.Request.Header.Get("X-AUTH-SECRET")
	rawUserID := ctx.Parameters.ByName("id")

	userId, err := strconv.ParseInt(rawUserID, 10, 64)
	if err != nil {
		c.badRequestResponse(ctx, errors.New("invalid user id provided"))
		return
	}

	user, err := c.ctx.Logic.GetUserByID(dto.Access{
		Token:  token,
		Secret: secret,
	}, userId)

	if err != nil {
		c.ServerErrorResponse(ctx, err)
		return
	}

	err = c.writeJSON(ctx, 200, user, nil)
	if err != nil {
		c.ServerErrorResponse(ctx, err)
	}
}

func (c *Controllers) BlockUserByID(ctx *WebContext) {
	token := ctx.Request.Header.Get("X-AUTH-TOKEN")
	secret := ctx.Request.Header.Get("X-AUTH-SECRET")
	rawUserID := ctx.Parameters.ByName("id")

	userID, err := strconv.ParseInt(rawUserID, 10, 64)
	if err != nil {
		c.badRequestResponse(ctx, errors.New("invalid user id provided"))
		return
	}

	user, err := c.ctx.Logic.BlockUserByID(dto.Access{
		Token:  token,
		Secret: secret,
	}, userID)

	if err != nil {
		c.ServerErrorResponse(ctx, err)
		return
	}

	err = c.writeJSON(ctx, 200, user, nil)
	if err != nil {
		c.ServerErrorResponse(ctx, err)
	}
}
