package logic

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	auth "github.com/dghubble/oauth1/twitter"

	"github.com/l1f/blockornot/internal/controller/dto"
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
		*accessToken.Token,
		*accessToken.Secret,
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

	requestTokenDTO := dto.Request{
		Url:    authUrl,
		Token:  &requestToken,
		Secret: &requestSecret,
	}

	return &requestTokenDTO, nil
}

func (l *logic) TwitterLoginResolve(requestToken dto.Request, pin string) (*dto.Access, error) {
	oauth1Config := l.getTwitterOAuthConfig()

	accessToken, accessSecret, err := oauth1Config.AccessToken(*requestToken.Token, *requestToken.Secret, pin)
	if err != nil {
		l.ctx.Logger.Error.Printf("couldn't get access requestToken: %v", err)
		return nil, err
	}

	accessTokenDTO := dto.Access{
		Token:  &accessToken,
		Secret: &accessSecret,
	}

	twitterClient := l.getAccountClient(accessTokenDTO)

	accountVerifyParams := &twitter.AccountVerifyParams{
		IncludeEntities: twitter.Bool(false),
		SkipStatus:      twitter.Bool(true),
		IncludeEmail:    twitter.Bool(false),
	}

	_, _, err = twitterClient.Accounts.VerifyCredentials(accountVerifyParams)
	if err != nil {
		l.ctx.Logger.Error.Printf("couldn't verify credentials: %v", err)
		return nil, err
	}

	return &accessTokenDTO, nil
}
