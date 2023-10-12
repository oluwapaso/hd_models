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

type PaymentList struct {
	Refrence_No         string `json:"refrence_no"`
	Status              string `json:"status"`
	Item_Type           string `json:"item_type"`
	Item_Id             string `json:"item_id"`
	Company_Id          string `json:"company_id"`
	Date_Received       string `json:"date_received"`
	Payment_From_To     string `json:"payment_from_to"`
	Amount              string `json:"amount"`
	Deposit_Method      string `json:"deposit_method"`
	Deposit_Type        string `json:"deposit_type"`
	CC_Last_Digit       string `json:"cc_last_digits"`
	CC_Type             string `json:"cc_type"`
	Other_CC_Type       string `json:"other_cc_type"`
	CC_Expiration_Month string `json:"cc_expiration_month"`
	CC_Expiration_Year  string `json:"cc_expiration_year"`
	Authorization_Code  string `json:"authorization_code"`
	Check_Number        string `json:"check_number"`
	Transaction_Id      string `json:"transaction_id"`
	Notes               string `json:"notes"`
	Receipt_Link        string `json:"receipt_link"`
	Entered_By          string `json:"entered_by"`
}
