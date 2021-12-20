package dto

import "net/url"

type Request struct {
	Url    *url.URL `json:"access_url"`
	Token  string   `json:"-"`
	Secret string   `json:"-"`
}

type Pin struct {
	Pin string `json:"pin"`
}

type Access struct {
	Token  string `json:"-"`
	Secret string `json:"-"`
}
