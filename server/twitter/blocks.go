package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/sling"
	"net/http"
)

// BlockService provides methods for accessing Twitter block endpoints
type BlockService struct {
	sling *sling.Sling
}

// newBlockService returns a new BlockService
func newBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{sling: sling.Path("blocks/")}
}

// BlockCreateParams are the parameters for BlockService.Create
type BlockCreateParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	SkipStatus      bool   `url:"skip_status,omitempty"`
}

// Create creates a new block for the authenticated user.
// https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/mute-block-report-users/api-reference/post-blocks-create
func (s *BlockService) Create(params *BlockCreateParams) (twitter.User, *http.Response, error) {
	user := new(twitter.User)
	apiError := new(twitter.APIError)
	resp, err := s.sling.Post("create.json").BodyForm(params).Receive(user, apiError)

	return *user, resp, relevantError(err, *apiError)
}

// BlockDestroyParams are the parameters for BlockService.Destroy.
type BlockDestroyParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"`
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

func (s *BlockService) Destroy(params *BlockDestroyParams) (twitter.User, *http.Response, error) {
	user := new(twitter.User)
	apiError := new(twitter.APIError)
	resp, err := s.sling.Post("destroy.json").BodyForm(params).Receive(user, apiError)

	return *user, resp, relevantError(err, *apiError)
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError twitter.APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}
