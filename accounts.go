package models

import "database/sql"

type Accounts struct {
	Account_id        int    `json:"account_id"`
	Company_unique_id string `json:"company_unique_id"`
	Created_by        string `json:"created_by"`
	Company_name      string `json:"company_name"`
	DefaultSettngs
}

type DefaultSettngs struct {
	Settings_id      int `json:"settings_id"`
	Company_id       int `json:"company_id"`
	Sendgrid_api_key int `json:"sendgrid_api_key"`
}

type AccountRepoResp struct {
	Resp RepoResponse
	Data []interface{}
}

type BrokerErrorInfo struct {
	Company_Id        int    `json:"settings_id"`
	Agent_id          int    `json:"agent_id"`
	Agent_name        string `json:"agent_name"`
	Category          string `json:"category"`
	Type              string `json:"type"`
	Complex_message   string `json:"complex_message"`
	Possible_error    string `json:"possible_error"`
	Possible_solution string `json:"possible_solution"`
	Affected_username string `json:"affected_username"`
	Simple_message    string `json:"simple_message"`
	Deal_type         string `json:"deal_type"` //Default None
	Deal_id           string `json:"deal_id"`
	Message_displayed string `json:"message_displayed"`
	Log_for_system    string `json:"log_for_system"` //Default No
	DBConn            *sql.DB
}

type ChargesDetails struct {
	Amount         int    `json:"amount"`
	Company_id     int    `json:"company_id"`
	Unique_ids     string `json:"unique_ids"`
	Table_name     string `json:"table_name"`
	Num_Of_Charges int    `json:"num_of_charges"`
}

type ValidateAPIKeyParams struct {
	Account_ID string
	API_Key    string
	APP_Stage  string
}

type ValidateWidgetKeyParams struct {
	Account_ID string
	Widget_Key string
	APP_Stage  string
}

type ValidateLeadSrcParams struct {
	Source_Email string
}
type ValidateLeadSrcKeyParams struct {
	Source_Key string
}

type ValidateAccntLevelParams struct {
	Agent_Id   int
	Company_Id int
}

type ReferralsListsParams struct {
	Type         string
	Company_Id   int
	Account_Type string
	Fields       string
	Order_By     string
	Paginated    string
	Start_From   int
	Limit        int
}

type CountReferralDealsParams struct {
	Referral_Id int
	LoadReportsParams
}

type ReferralsDispDataParams struct {
	Referral_Id int
	LoadReportsParams
}

type ReferralsReportData struct {
	Referral_Name    string  `json:"referral"`
	Status           string  `json:"status"`
	Total_Leads_Amnt float64 `json:"total_leads_amount"`
	DealsReportData
}

type UpdateDlftRoutingNumParams struct {
	Company_Id   int
	Phone_Number string
	Tx           *sql.Tx
}

type LoadAccountInfoParams struct {
	Load_By    string
	Email      string
	Account_Id int
	Fields     string
}

type UpdateLogoParams struct {
	Company_Id int
	Logo       string
	Tx         *sql.Tx
}

type UpdateExtBGParams struct {
	Company_Id int
	BG_Loc     string
	Tx         *sql.Tx
}

type UpdateChargesParams struct {
	Company_Id     int
	New_Credit_Bal int
	Tx             *sql.Tx
}

type AccntWithTwilioParams struct {
	To_Number string
	Fields    string
}

type ListBrokersParams struct {
	Type      string
	Fields    string
	Status    string
	Sub_Type  string
	Keyword   string
	StartFrom int
	Limit     int
}
