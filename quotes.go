package models

import "database/sql"

type Quotes struct {
	QuoteFields
	CommonDealFields
}

type QuoteResp struct {
	RepoRes RepoResponse
	Data    Quotes
}

type QuotesListsParams struct {
	Type       string
	Company_Id int
	Agent_Id   string
	Fields     string
	Order_By   string
	Paginated  string
	Start_From int
	Limit      int
	LoadReportsParams
}

type QuoteLists struct {
	Quote_Id             int         `json:"quote_id"`
	Quote_Uniq_Id        string      `json:"quote_uniq_id"`
	Company_Id           int         `json:"company_id"`
	Dates                interface{} `json:"dates"`
	Car_Run              string      `json:"car_run"`
	Ship_Via             string      `json:"ship_via"`
	Assigned_To          int         `json:"assigned_to"`
	Original_Assigned_To int         `json:"original_assigned_to"`
	Assigned_To_Info     interface{} `json:"assigned_to_info"`
	Origin               interface{} `json:"origin"`
	Destination          interface{} `json:"destination"`
	Vehicles             interface{} `json:"vehicles"`
	Shipper              interface{} `json:"shipper"`
	Payments             interface{} `json:"payments"`
	Type                 string      `json:"type"`
	Source               string      `json:"source"`
	Source_Type          string      `json:"source_type"`
	Status               string      `json:"status"`
	Item_Counts          interface{} `json:"item_counts"`
	Subscribed_Drips     interface{} `json:"subscribed_drips"`
}

type QuoteDetails struct {
	QuoteLists
	//Not finished yet 11/22/2023
}

type UpdateQuoteParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
	Must_Update  string
}

type UpdateQuoteFiledsParams struct {
	Tx              *sql.Tx
	Company_Id      int
	User_Id         int
	View_User       string
	Quote_Id        int
	Quote_Uniq_Id   string
	Referral        string
	Shipper         string
	Pickup          string
	Drop_Off        string
	Car_Run         string
	Ship_Via        string
	No_Notes_From   int //use json_set
	Changed_Details string
	Called_From     string
	Agent_Name      string
	Resources       string

	/** Update For Internal Only Starts **/
	Dates      string //Has move_on //So use use json_set
	Load_On    string
	Deliver_On string
	Tags       string
	Payments   string
	/** Update For Internal Only Ends **/

	Vehicles             string
	Vehicle_Info_Changed string
	Tariff_Changed       string
	Old_Tariff           string
	New_Tariff           string
}
