package controller

import "github.com/l1f/blockornot/internal/application"

type Controllers struct {
	ctx *application.Context
}

func New(ctx *application.Context) Controllers {
	return Controllers{ctx: ctx}
}
