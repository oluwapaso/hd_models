package models

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
	Type         string
	Fields       string
	InitQuery    string
	AssignedTo   string
	TagQuery     string
	SearchQuery  string
	Params       []interface{}
	TagParams    []interface{}
	SearchParams []interface{}
	StartFrom    int
	Limit        int
}

// Used for dynamic fetching so sql.NullString and likes will not be used here
type OrderLists struct {
	Order_Id          int    `json:"order_id"`
	Order_Uniq_Id     string `json:"order_uniq_id"`
	Company_Id        int    `json:"company_id"`
	Dates             string `json:"dates"`
	Car_Run           string `json:"car_run"`
	Ship_Via          string `json:"ship_via"`
	Assigned_To       int    `json:"assigned_to"`
	Assigned_To_Info  string `json:"assigned_to_info"`
	Origin            string `json:"origin"`
	Destination       string `json:"destination"`
	Dispatch_Details  string `json:"dispatch_details"`
	Vehicles          string `json:"vehicles"`
	Shipper           string `json:"shipper"`
	Payments          string `json:"payments"`
	Type              string `json:"type"`
	Source            string `json:"source"`
	Source_Type       string `json:"source_type"`
	Status            string `json:"status"`
	Item_Counts       string `json:"item_counts"`
	Issues            string `json:"issues"`
	Shipper_Signature string `json:"shipper_signature"`
	Subscribed_Drips  string `json:"subscribed_drips"`
}
type OrderDetails struct {
	OrderLists
	Mode                 string `json:"mode,omitempty"`
	Original_Assigned_To string `json:"original_assigned_to,omitempty"`
	Source_Id            string `json:"source_id,omitempty"`
	Referral             string `json:"referral,omitempty"`
	Load_On              string `json:"load_on,omitempty"`
	Deliver_On           string `json:"deliver_on,omitempty"`
	Posted_Board         string `json:"posted_board,omitempty"`
	Original_Email       string `json:"original_email,omitempty"`
	HDesk_Post_Id        string `json:"hdesk_post_id,omitempty"`
	Resolve_Issue_To     string `json:"resolve_issue_to,omitempty"`
	Quote_Id_Converted   string `json:"quote_id_converted,omitempty"`
	Quote_Uniq_Id_Conv   string `json:"quote_uniq_id_conv,omitempty"`
	Tracking_Number      string `json:"tracking_number,omitempty"`
	Converted_Details    string `json:"converted_details,omitempty"`
	Order_Token          string `json:"order_token,omitempty"`
	CD_Dispatch_Id       string `json:"cd_dispatch_id,omitempty"`
	CD_Dispatch_Status   string `json:"cd_dispatch_status,omitempty"`
	User_Order_Id        string `json:"user_order_id,omitempty"`
	Automate             string `json:"automate,omitempty"`
	Driver_Id            string `json:"driver_id,omitempty"`
	Carrier_Id           string `json:"carrier_id,omitempty"`
	Driver_Signature     string `json:"driver_signature,omitempty"`
	Automation_Id        string `json:"automation_id,omitempty"`
	Automation_Status    string `json:"automation_status,omitempty"`
	Tags                 string `json:"tags,omitempty"`
	Dates_Created        string `json:"dates_created,omitempty"`
	Shpr_Shipper_Id      string `json:"shpr_shipper_id,omitempty"`
	Dispatched_To        string `json:"dispatched_to,omitempty"`
	P_Terminal_Id        string `json:"p_terminal_id,omitempty"`
	D_Terminal_Id        string `json:"d_terminal_id,omitempty"`
	Referral_Id          string `json:"referral_id,omitempty"`
}

type ListOrdersResponse struct {
	DB_Query_Response
	Orders  []Orders
	Message string
}
