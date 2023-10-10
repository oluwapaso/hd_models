package models

import "database/sql"

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

type UpdatePaymentsParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}

type Payments struct {
	Payment_Id         int            `json:"payment_id"`
	Refrence_No        string         `json:"refrence_no"`
	Order_Id           int            `json:"order_id"`
	Company_Id         string         `json:"company_id"`
	Date_Received      string         `json:"date_received"`
	Payment_From_To    string         `json:"payment_from_to"`
	Amount             string         `json:"amount"`
	Deposit_Method     sql.NullString `json:"deposit_method"`
	Authorization_Code sql.NullString `json:"authorization_code"`
	Transaction_Id     sql.NullString `json:"transaction_id"`
	Notes              sql.NullString `json:"notes"`
	Entered_By         sql.NullString `json:"entered_by"`
}

type AddPaymentsParams struct {
	Refrence_No     string
	Order_Id        int
	Order_Unique_Id string
	Company_Id      string
	Date_Received   string
	Payment_From_To string
	Amount          float64
	Deposit_Method  string
	Transaction_Id  string
	Notes           string
	Entered_By      string
}
