package models

import "database/sql"

type UpdateAdminStatusParams struct {
	Update_By   string
	Admin_Id    int
	Email       string
	Status      string
	Must_Update bool
	Tx          *sql.Tx
}

type UpdateAdminPassParams struct {
	Password string
	Admin_Id int
	Tx       *sql.Tx
}

type AddNewAdminSessionParams struct {
	Admin_Id            int
	Access_Token        string
	Access_Token_expiry string
	Last_Login          string
	Tx                  *sql.Tx
}

type DeleteAdminSessionsParams struct {
	Delete_By  string
	Admin_Id   int
	Session_Id int
	Tx         *sql.Tx
}

type ValidateAdminTokenParams struct {
	Admin_Id     int
	Access_Token string
}

type AdminTokenDetails struct {
	ID                  string `json:"id"`
	Access_Token_Expiry string `json:"access_token_expiry"`
}

type AdminDashboardData struct {
	PaidAccounts          int `json:"paid_accounts"`
	FreeAccounts          int `json:"free_accounts"`
	ActiveAccounts        int `json:"active_accounts"`
	InactiveAccounts      int `json:"inactive_accounts"`
	ActiveSubscriptions   int `json:"active_subscriptions"`
	InactiveSubscriptions int `json:"inactive_subscriptions"`
}

type SystemAlertsParams struct {
	Search_Type string
	Fields      string
	StartFrom   int
	Limit       int
	Type        string
}
