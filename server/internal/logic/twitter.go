package logic

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	auth "github.com/dghubble/oauth1/twitter"

	"github.com/l1f/blockornot/internal/controller/dto"
	"github.com/l1f/blockornot/internal/logic/types"
)

func (l *logic) getTwitterOAuthConfig() *oauth1.Config {
	return &oauth1.Config{
		ConsumerKey:    l.ctx.Config.Twitter.ConsumerKey,
		ConsumerSecret: l.ctx.Config.Twitter.ConsumerSecret,
		CallbackURL:    "oob",
		Endpoint:       auth.AuthenticateEndpoint,
	}
}

func (l *logic) getAccountClient(accessToken dto.Access) *twitter.Client {
	config := oauth1.NewConfig(
		l.ctx.Config.Twitter.ConsumerKey,
		l.ctx.Config.Twitter.ConsumerSecret,
	)
	token := oauth1.NewToken(
		accessToken.Token,
		accessToken.Secret,
	)

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}

func (l *logic) TwitterLoginInit() (*dto.Request, error) {
	oauth1Config := l.getTwitterOAuthConfig()
	requestToken, requestSecret, err := oauth1Config.RequestToken()
	if err != nil {
		l.ctx.Logger.Error.Printf("couldn't get requestToken: %v", err)
		return nil, err
	}

	authUrl, err := oauth1Config.AuthorizationURL(requestToken)
	if err != nil {
		l.ctx.Logger.Error.Printf("couldn't get authorization url: %v", err)
		return nil, err
	}

	return &dto.Request{
		Url:    authUrl,
		Token:  requestToken,
		Secret: requestSecret,
	}, nil
}

func (l *logic) TwitterLoginResolve(requestToken dto.Request, pin string) (*dto.Access, *dto.Account, error) {
	oauth1Config := l.getTwitterOAuthConfig()

	accessToken, accessSecret, err := oauth1Config.AccessToken(requestToken.Token, requestToken.Secret, pin)
	if err != nil {
		l.ctx.Logger.Error.Printf("couldn't get access requestToken: %v", err)
		return nil, nil, err
	}

	accessTokenDTO := dto.Access{
		Token:  accessToken,
		Secret: accessSecret,
	}

	twitterClient := l.getAccountClient(accessTokenDTO)

	accountVerifyParams := &twitter.AccountVerifyParams{
		IncludeEntities: twitter.Bool(false),
		SkipStatus:      twitter.Bool(true),
		IncludeEmail:    twitter.Bool(false),
	}

	user, _, err := twitterClient.Accounts.VerifyCredentials(accountVerifyParams)
	if err != nil {
		l.ctx.Logger.Error.Printf("couldn't verify credentials: %v", err)
		return nil, nil, err
	}

	return &accessTokenDTO, &dto.Account{
		ScreenName: user.ScreenName,
		Name:       user.Name,
		TwitterID:  user.ID,
		AvatarURL:  user.ProfileImageURL,
	}, nil
}

func (l logic) SearchTweets(tokens dto.Access, query string, result *types.TwitterSearchResultType) (*[]dto.Tweet, error) {
	client := l.getAccountClient(tokens)

	defaultResultType := types.TwitterSearchResultPopular
	if result != nil {
		defaultResultType = *result
	}

	tweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:           query,
		ResultType:      string(defaultResultType),
		Count:           15,
		IncludeEntities: nil,
		// https://github.com/tweepy/tweepy/issues/1170
		TweetMode: "extended",
	})

	if err != nil {
		return nil, err
	}

	var tweetDto []dto.Tweet
	for _, t := range tweets.Statuses {
		tweetDto = append(tweetDto, dto.Tweet{
			CreatedAt:         t.CreatedAt,
			Entities:          t.Entities,
			FavoriteCount:     t.FavoriteCount,
			Favorited:         t.Favorited,
			ID:                t.ID,
			PossiblySensitive: t.PossiblySensitive,
			QuoteCount:        t.QuoteCount,
			ReplyCount:        t.ReplyCount,
			RetweetCount:      t.RetweetCount,
			Retweeted:         t.Retweeted,
			RetweetedStatus:   t.RetweetedStatus,
			Text:              t.Text,
			FullText:          t.FullText,
			User:              t.User,
		})
	}

	return &tweetDto, nil
}
