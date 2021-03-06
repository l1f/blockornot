package application

import (
	"github.com/l1f/blockornot/internal/controller/dto"
	"github.com/l1f/blockornot/internal/logic/types"
)

type Logic interface {
	TwitterLoginInit() (*dto.Request, error)
	TwitterLoginResolve(requestToken dto.Request, pin string) (*dto.Access, *dto.Account, error)

	SearchTweets(tokens dto.Access, query string, result *types.TwitterSearchResultType) (*[]dto.Tweet, error)
	GetUserByID(tokens dto.Access, userID int64) (*dto.Account, error)
	BlockUserByID(tokens dto.Access, userId int64) (*dto.Account, error)
}
