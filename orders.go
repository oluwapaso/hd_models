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
