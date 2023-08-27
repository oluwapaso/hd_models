package models

type Quotes struct {
	QuoteFields
	CommonDealFields
}

type QuoteResp struct {
	RepoRes RepoResponse
	Data    Quotes
}
