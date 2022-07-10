package twitter

import (
	baseClient "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/sling"
	"net/http"
)

const twitterAPI = "https://api.twitter.com/1.1/"

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	Client *baseClient.Client
	Blocks *BlockService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPI)

	client := Client{
		Client: baseClient.NewClient(httpClient),
		Blocks: newBlockService(base),
	}

	return &client
}
