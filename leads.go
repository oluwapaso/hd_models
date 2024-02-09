package models

import (
	"database/sql"
)

type LeadLists struct {
	Lead_Id              int         `json:"lead_id"`
	Lead_Uniq_Id         string      `json:"lead_uniq_id"`
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

type LeadDetails struct {
	LeadLists
	Original_Assigned_To string      `json:"original_assigned_to,omitempty"`
	Source_Id            string      `json:"source_id,omitempty"`
	Referral             string      `json:"referral,omitempty"`
	Tags                 interface{} `json:"tags,omitempty"`
	Dates_Received       string      `json:"received,omitempty"`
	Referral_Id          string      `json:"referral_id,omitempty"`
	//Not finished yet 11/22/2023
}

type ListLeadsParams struct {
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
	Curr_Lead_Id    string
	Filter_Field    string
	Filter_Val      string
	Next_Filter     string
	Previous_Filter string
	Sort_Dir        string
}

type UpdateLeadStatusParams struct {
	Company_Id int
	Lead_Ids   string
	Status     string
	Tx         *sql.Tx
}

type ReassignLeadsParams struct {
	Company_Id  int
	Lead_Ids    string
	Reassign_To int
	Tx          *sql.Tx
}

type UpdateLeadFieldsParams struct {
	Tx              *sql.Tx
	Company_Id      int
	User_Id         int
	View_User       string
	Lead_Id         int
	Lead_Uniq_Id    string
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
	Dates    string //Has move_on, load_date, delivery_date //So use use json_set
	Tags     string
	Payments string
	/** Update For Internal Only Ends **/

	Vehicles             string
	Vehicle_Info_Changed string
	Tariff_Changed       string
	Old_Tariff           string
	New_Tariff           string
}

type UpdateLeadParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
	Must_Update  string
}

type ConvertToQuoteParams struct {
	Company_Id      int
	Company_Info    []interface{}
	Agent_Id        int
	Lead_Ids        []interface{}
	Lead_Unique_Ids []interface{}
	Send_Email      string
	Temp_Info       []interface{}
	Tx              *sql.Tx
}

type LeadConvResp struct {
	Converted           bool
	Error_Msg           string
	Last_Quote_Id       int
	All_Quote_Id        []interface{}
	All_Quote_Ass_To    []interface{}
	Last_Assigned_Uname string
}

type CheckDuplicateParams struct {
	Check_Type    string
	Vehicle       string
	Company_Id    int
	Shipper_Email string
	Tx            *sql.Tx
}

type AddNewLeadParams struct {
	Tx          *sql.Tx
	Company_Id  int
	User_Id     int
	View_User   string
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
	Payments string
	/** Update For Internal Only Ends **/

	Source_Id      int
	Source_Name    string
	Source_Type    string
	Lead_Medium    string
	Duplicate_Info string
	Original_Email string
	Seen_By_Agent  string
}

type UpdateNewLeadParams struct {
	Company_Id   int
	User_Id      int
	Lead_Id      int
	Lead_Uniq_Id string
	Item_Token   string
	Lead_Type    string
	Payments     string
	Tx           *sql.Tx
}
