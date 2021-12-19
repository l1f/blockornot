package service

import "github.com/l1f/blockornot/internal/application"

type Service struct {
	ctx *application.Context
}

func New(ctx *application.Context) application.ServiceInterface {
	return Service{ctx: ctx}
}
