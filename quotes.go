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
	Source_Id       string      `json:"source_id,omitempty"`
	Referral        string      `json:"referral,omitempty"`
	Original_Email  string      `json:"original_email,omitempty"`
	Quote_Token     string      `json:"quote_token,omitempty"`
	Tags            interface{} `json:"tags,omitempty"`
	Dates_Created   string      `json:"dates_created,omitempty"`
	Shpr_Shipper_Id string      `json:"shpr_shipper_id,omitempty"`
	Referral_Id     string      `json:"referral_id,omitempty"`
	//Not finished yet 11/22/2023
}

type UpdateQuoteParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
	Must_Update  string
}

type UpdateQuoteFieldsParams struct {
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

type ListQuotesParams struct {
	Type            string
	Fields          string
	InitQuery       string
	AssignedTo      string
	TagQuery        string
	SearchQuery     string
	Params          []interface{}
	TagParams       []interface{}
	SearchParams    []interface{}
	StartFrom       int
	Limit           int
	Order_By        string
	Curr_Quote_Id   string
	Filter_Field    string
	Filter_Val      string
	Next_Filter     string
	Previous_Filter string
	Sort_Dir        string
}

type ReassignQuotesParams struct {
	Company_Id  int
	Quote_Ids   string
	Reassign_To int
	Tx          *sql.Tx
}

type UpdateQuoteStatusParams struct {
	Company_Id int
	Quote_Ids  string
	Status     string
	Tx         *sql.Tx
}

type UpdateQuoteExpiryDateParams struct {
	Company_Id int
	Quote_Ids  string
	Expire_On  string
	Tx         *sql.Tx
}

type UpdateFupedQtsParams struct {
	Company_Id         int
	Quote_Ids          string
	Next_Followup_Date string
	Last_Bomb          string
	Tx                 *sql.Tx
}

type UpdateQtFupInfoParams struct {
	Company_Id int
	Quote_Id   int
	DbCon      *sql.DB
}

type AddNewQuoteParams struct {
	Tx          *sql.Tx
	Company_Id  int
	User_Id     int
	View_User   string
	Referral    string
	Shipper     string
	Pickup      string
	Drop_Off    string
	Car_Run     string
	Ship_Via    string
	Item_Counts string
	Agent_Name  string
	Vehicles    string

	/** Update For Internal Only Starts **/
	Dates    string //Has move_on //So use use json_set
	Tags     string
	Payments string
	/** Update For Internal Only Ends **/

	Broker_Comp_Name  string
	Broker_Comp_Email string
	Source_Id         int
	Source_Name       string
	Source_Type       string
	Lead_Medium       string
	Duplicate_Info    string
	Follow_Up_Info    string
	Original_Email    string
	Seen_By_Agent     string
}

type UpdateNewQuoteParams struct {
	Company_Id    int
	User_Id       int
	Quote_Id      int
	Quote_Uniq_Id string
	Item_Token    string
	Payments      string
	Tx            *sql.Tx
}
