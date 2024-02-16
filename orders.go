package models

import "database/sql"

type Orders struct {
	OrderFields
	CommonDealFields
}

type OrderResp struct {
	RepoRes RepoResponse
	Data    Orders
}

type SyncCDOrderParams struct {
	Company_Id          string `json:"company_id"`
	Order_Id            int    `json:"order_id"`
	Order_Uniq_Id       string `json:"order_uniq_id"`
	Assigned_To         string `json:"assigned_to"`
	Created_By          string `json:"created_by"`
	Order_Type          string `json:"order_type"`
	Current_Status      string `json:"current_status"`
	No_Internal_Notes   string `json:"no_internal_notes"`
	CD_Dispatch_Id      string `json:"cd_dispatch_id"`
	CD_Dispatch_Status  string `json:"cd_dispatch_status"`
	User_Order_Id       string `json:"user_order_id"`
	Last_CD_Status_Sync string `json:"last_cd_status_sync"`
	CD_username         string `json:"cd_username"`
	CD_password         string `json:"cd_password"`
	Time_Zone           string `json:"time_zone"`
	Cookies             string `json:"cookies"`
}

type ListOrdersParams struct {
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
	Curr_Order_Id   string
	Filter_Field    string
	Filter_Val      string
	Next_Filter     string
	Previous_Filter string
	Sort_Dir        string
}

// Used for dynamic fetching so sql.NullString and likes will not be used here
type OrderLists struct {
	Order_Id             int         `json:"order_id"`
	Order_Uniq_Id        string      `json:"order_uniq_id"`
	Company_Id           int         `json:"company_id"`
	Dates                interface{} `json:"dates"`
	Car_Run              string      `json:"car_run"`
	Ship_Via             string      `json:"ship_via"`
	Assigned_To          int         `json:"assigned_to"`
	Original_Assigned_To int         `json:"original_assigned_to"`
	Assigned_To_Info     interface{} `json:"assigned_to_info"`
	Origin               interface{} `json:"origin"`
	Destination          interface{} `json:"destination"`
	Dispatch_Details     interface{} `json:"dispatch_details"`
	Vehicles             interface{} `json:"vehicles"`
	Shipper              interface{} `json:"shipper"`
	Payments             interface{} `json:"payments"`
	Type                 string      `json:"type"`
	Source               string      `json:"source"`
	Source_Type          string      `json:"source_type"`
	Status               string      `json:"status"`
	Item_Counts          interface{} `json:"item_counts"`
	Issues               interface{} `json:"issues"`
	Shipper_Signature    string      `json:"shipper_signature"`
	Subscribed_Drips     interface{} `json:"subscribed_drips"`
}
type OrderDetails struct {
	OrderLists
	Mode                 string      `json:"mode,omitempty"`
	Original_Assigned_To string      `json:"original_assigned_to,omitempty"`
	Source_Id            string      `json:"source_id,omitempty"`
	Referral             string      `json:"referral,omitempty"`
	Load_On              string      `json:"load_on,omitempty"`
	Deliver_On           string      `json:"deliver_on,omitempty"`
	Posted_Board         interface{} `json:"posted_board,omitempty"`
	Original_Email       string      `json:"original_email,omitempty"`
	HD_Post_Id           string      `json:"hd_post_id,omitempty"`
	Resolve_Issue_To     string      `json:"resolve_issue_to,omitempty"`
	Quote_Id_Converted   string      `json:"quote_id_converted,omitempty"`
	Quote_Uniq_Id_Conv   string      `json:"quote_uniq_id_conv,omitempty"`
	Tracking_Number      string      `json:"tracking_number,omitempty"`
	Converted_Details    string      `json:"converted_details,omitempty"`
	Order_Token          string      `json:"order_token,omitempty"`
	CD_Dispatch_Id       string      `json:"cd_dispatch_id,omitempty"`
	CD_Dispatch_Status   string      `json:"cd_dispatch_status,omitempty"`
	User_Order_Id        string      `json:"user_order_id,omitempty"`
	Automate             string      `json:"automate,omitempty"`
	Driver_Id            string      `json:"driver_id,omitempty"`
	Carrier_Id           string      `json:"carrier_id,omitempty"`
	Driver_Signature     string      `json:"driver_signature,omitempty"`
	Automation_Id        string      `json:"automation_id,omitempty"`
	Automation_Status    string      `json:"automation_status,omitempty"`
	Tags                 interface{} `json:"tags,omitempty"`
	Dates_Created        string      `json:"dates_created,omitempty"`
	Shpr_Shipper_Id      string      `json:"shpr_shipper_id,omitempty"`
	Dispatched_To        string      `json:"dispatched_to,omitempty"`
	P_Terminal_Id        string      `json:"p_terminal_id,omitempty"`
	D_Terminal_Id        string      `json:"d_terminal_id,omitempty"`
	Referral_Id          string      `json:"referral_id,omitempty"`
}

type ListOrdersResponse struct {
	DB_Query_Response
	Orders  []Orders
	Message string
}

type InsertOrderParams struct {
	Tx           *sql.Tx
	InsertField  string
	InsertValues []interface{}
}

type CollectSignatureParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}

type UpdateOrderParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
	Must_Update  string
}

type OrderIssue struct {
	Id       string `json:"id"`
	Date     string `json:"date"`
	Type     string `json:"type"`
	Notes    string `json:"notes"`
	Added_By string `json:"added_by"`
}

type UpdateOrderFieldsParams struct {
	Tx              *sql.Tx
	Company_Id      int
	User_Id         int
	View_User       string
	Order_Id        int
	Order_Uniq_Id   string
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
	Dates            string //Has move_on, load_date, delivery_date //So use use json_set
	Load_On          string
	Deliver_On       string
	Tags             string
	Dispatch_Details string
	Payments         string
	/** Update For Internal Only Ends **/

	Vehicles             string
	Vehicle_Info_Changed string
	Tariff_Changed       string
	Old_Tariff           string
	New_Tariff           string
}

type AddNewOrderParams struct {
	Tx                   *sql.Tx
	Company_Id           int
	User_Id              int
	Assigned_To          int
	Assigned_To_Info     string
	Original_Assigned_To int
	View_User            string
	Referral             string
	Shipper              string
	Pickup               string
	Drop_Off             string
	Car_Run              string
	Ship_Via             string
	Item_Counts          string
	Agent_Name           string
	Vehicles             string
	Tracking_Number      string

	/** Update For Internal Only Starts **/
	Dates            string //Has move_on //So use use json_set
	Tags             string
	Payments         string
	Dispatch_Details string
	/** Update For Internal Only Ends **/

	Lead_Medium    string
	Source_Id      int
	Source_Name    string
	Source_Type    string
	Original_Email string
	Quote_Id_Conv  string
	Quote_UId_Conv string
	Seen_By_Agent  string
	Mode           string
	Load_On        string
	Deliver_On     string
}

type UpdateNewOrderParams struct {
	Company_Id    int
	User_Id       int
	Order_Id      int
	Order_Uniq_Id string
	Item_Token    string
	Payments      string
	Tx            *sql.Tx
}

type ReassignOrdersParams struct {
	Company_Id  int
	Order_Ids   string
	Reassign_To int
	Tx          *sql.Tx
}

type UpdateOrderStatusParams struct {
	Company_Id int
	Order_Ids  string
	Status     string
	Tx         *sql.Tx
}

type MarkAsIssueParams struct {
	Company_Id       int
	Order_Id         int
	Id               int64
	Issues_Type      string
	Date             string
	Internal_Notes   string
	Added_By         string
	Issue            string
	Resolve_Issue_To string
	Tx               *sql.Tx
}

type ResolveIssueParams struct {
	Company_Id int
	Status     string
	Issues     string
	Order_Ids  string
	Tx         *sql.Tx
}

type UpdateReservationParams struct {
	Reservation_Type string
	Company_Id       int
	Order_Id         int
	Tx               *sql.Tx
}

type MonitorabaleOrdersParams struct {
	Fields string
	Now    string
	Limit  int
}
