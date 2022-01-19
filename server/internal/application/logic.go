package application

import (
	"github.com/dghubble/go-twitter/twitter"

	"github.com/l1f/blockornot/internal/controller/dto"
	"github.com/l1f/blockornot/internal/logic/types"
)

type Logic interface {
	TwitterLoginInit() (*dto.Request, error)
	TwitterLoginResolve(requestToken dto.Request, pin string) (*dto.Access, *dto.Account, error)
	SearchTweets(tokens dto.Access, query string, result *types.ResultType) (*[]twitter.Tweet, error)
}
