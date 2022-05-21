package dto

import "github.com/dghubble/go-twitter/twitter"

type Tweet struct {
	CreatedAt         string            `json:"created_at"`
	Entities          *twitter.Entities `json:"entities"`
	FavoriteCount     int               `json:"favorite_count"`
	Favorited         bool              `json:"favorited"`
	ID                int64             `json:"id"`
	PossiblySensitive bool              `json:"possibly_sensitive"`
	QuoteCount        int               `json:"quote_count"`
	ReplyCount        int               `json:"reply_count"`
	RetweetCount      int               `json:"retweet_count"`
	Retweeted         bool              `json:"retweeted"`
	RetweetedStatus   *twitter.Tweet    `json:"retweeted_status"`
	Text              string            `json:"text"`
	FullText          string            `json:"full_text"`
	User              *twitter.User     `json:"user"`
}
