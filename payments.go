package models

type OpenBalanceParams struct {
	Company_Id    int
	Item_Type     string
	Item_Id       int
	Deal_Payments interface{}
}

type LoadPaymentsParams struct {
	Search_Type string
	Company_Id  int
	Item_Type   string
	Item_Id     int
	Fields      string
}
