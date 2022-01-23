package types

type TwitterSearchResultType string

const (
	TwitterSearchResultMixed   TwitterSearchResultType = "mixed"
	TwitterSearchResultRecent  TwitterSearchResultType = "recent"
	TwitterSearchResultPopular TwitterSearchResultType = "popular"
)
