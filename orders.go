package models

type Orders struct {
	OrderFields
	CommonDealFields
}

type OrderResp struct {
	RepoRes RepoResponse
	Data    Orders
}
