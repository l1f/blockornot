package application

import "github.com/l1f/blockornot/internal/controller/dto"

type Logic interface {
	TwitterLoginInit() (*dto.Request, error)
	TwitterLoginResolve(requestToken dto.Request, pin string) (*dto.Access, *dto.Account, error)
}
