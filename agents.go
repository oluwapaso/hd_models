package models

import "database/sql"

type Agents struct {
	Agent_id         int            `json:"agent_id"`
	Company_id       string         `json:"company_id"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone,omitempty"`
	Lead_multiples   int            `json:"lead_multiples,omitempty"`
	Username         string         `json:"username,omitempty"`
	Mapped_import    string         `json:"mapped_import,omitempty"`
	Webmail_settings sql.NullString `json:"webmail_settings,omitempty"`
	Mailer_settings  sql.NullString `json:"mailer_settings,omitempty"`
}

type MultiAgentsResp struct {
	RepoRes RepoResponse
	Data    []Agents
}

type ValidateTokenParams struct {
	Company_Id   int
	Agent_Id     int
	Access_Token string
}

type TokenDetails struct {
	ID                  string `json:"id"`
	Access_Token_Expiry string `json:"access_token_expiry"`
	Account_Status      string `json:"account_status"`
	Subscription_Info   string `json:"subscription_info"`
}

type ValidateSecurityParams struct {
	User_Id    int    `json:"user_id"`
	Company_Id int    `json:"company_id"`
	Field      string `json:"field"`
}

type UpdateAgentSortParams struct {
	Company_Id     int
	Agent_Id       int
	Update_Type    string
	Direction      string
	Sort_By        string
	Update_Literal string
}

type AgentsListsParams struct {
	Type       string
	Company_Id int
	Agent_Id   string
	Fields     string
	Order_By   string
	Paginated  string
	Start_From int
	Limit      int
}

type CountAgentsDealsParams struct {
	Agent_Id int
	LoadReportsParams
}

type AgentsDispDataParams struct {
	Agent_Id         int
	Define_Orders_As string
	LoadReportsParams
}

type AgentsReportData struct {
	Agent_Id        int     `json:"agent_id"`
	Agent_Name      string  `json:"agent_name"`
	Agent_Status    string  `json:"status"`
	Total_Comm_Amnt float64 `json:"total_comm_amount"`
	DealsReportData
}

type UpdateAgentRoutingNumParams struct {
	Company_Id   int
	Agent_Id     int
	Phone_Number string
	Tx           *sql.Tx
}
