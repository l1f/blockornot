package dto

import "net/url"

type Request struct {
	Url    *url.URL `json:"access_url"`
	Token  *string  `json:"request_token"`
	Secret *string  `json:"request_token_secret"`
}

type Access struct {
	Token  *string `json:"access_token"`
	Secret *string `json:"access_token_secret"`
}
