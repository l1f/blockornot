package logic

import "github.com/l1f/blockornot/internal/application"

type logic struct {
	ctx *application.Context
}

func New(ctx *application.Context) application.Logic {
	return &logic{ctx: ctx}
}
