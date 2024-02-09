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
	WebmailSettings  string         `json:"webmail_sett,omitempty"`
	MailerSettings   string         `json:"mailer_sett,omitempty"`
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

type ValidateAssItemsParams struct {
	Agent_Id             int
	Company_Id           int
	Item_Array           []interface{}
	Item_Type            interface{}
	Required_Permissions []interface{}
}

type ValMultiSecParams struct {
	Agent_Id    int
	Company_Id  int
	Permissions []interface{}
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
	Status     string
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

type UpdateAgentMailer struct {
	Update_Type    string
	Company_Id     int
	Default_Mailer interface{}
	Allow_SMTP     string
	Noti_Sett      string
	Agent_Id       int
	Tx             *sql.Tx
}

type UpdateAgentWebMailer struct {
	Company_Id         int
	Webmailer_Settings interface{}
	Agent_Id           int
	Tx                 *sql.Tx
}

type UpdateAgentAccntGrp struct {
	Company_Id     int
	Update_Type    string
	Group_Name     string
	Old_Group_Name string
	Agent_Id       int
	Permissions    string
	Tx             *sql.Tx
}

type LoadAgentInfoParams struct {
	Load_By  string
	Email    string
	Username string
	Agent_Id int
	Fields   string
}

type AddNewAgentParams struct {
	Company_Id             int
	Fullname               string
	Email                  string
	Phone                  string
	Username               string
	Password               string
	Permissions            string
	Deal_List_Settings     string
	Notifications_Settings string
	Date                   string
	User_Mailer            string
	Status                 string
	Level                  string
	Role                   string
	Tx                     *sql.Tx
}

type LoginHistParams struct {
	Company_Id int
	Agent_Id   int
	Start_Date string
	End_Date   string
	Fields     string
	Limit      int
}

type UpdateAgentNotiParams struct {
	Agent_Id               int
	Company_Id             int
	Notifications_Settings string
	Tx                     *sql.Tx
}

type UpdateAgentInfoParams struct {
	Agent_Id            int
	Company_Id          int
	Fullname            string
	Email               string
	Phone               string
	Status              string
	Role                string
	Commission_Settings string
	Tx                  *sql.Tx
}

type UpdateAgentLoginInfoParams struct {
	Agent_Id   int
	Company_Id int
	Username   string
	Password   string
	Tx         *sql.Tx
}

type UpdateDPParams struct {
	Company_Id int
	Agent_Id   int
	DP_Loc     string
	Tx         *sql.Tx
}

type UpdateAgentStatusParams struct {
	Update_By   string
	Agent_Id    int
	Email       string
	Status      string
	Must_Update bool
	Tx          *sql.Tx
}

type UpdateAgentPassParams struct {
	Password string
	Agent_Id int
	Tx       *sql.Tx
}

type UpdateLastLoginParams struct {
	Agent_Id   int
	Last_Login string
	Tx         *sql.Tx
}

type GetAgentFor_RR_Params struct {
	Company_Id    int
	Fields        string
	Search_Type   string
	Exclude_Agent int
	Group         string
}

type Agents_SRC_Params struct {
	Company_Id     int
	Fields         string
	Start_From     int
	Limit          int
	Lead_Source_Id int
}

type Agent_SRC_Sett_Params struct {
	Company_Id     int
	Lead_Source_Id int
	Agent_Id       int
}

type AddAgentPaymentsParams struct {
	Company_Id         int
	Agent_Id           int
	Date               string
	Amount             string
	Payment_Type       string
	Payment_Method     string
	Transaction_Number string
	Notes              string
	Tx                 *sql.Tx
}

type AgentsPaymentsListsParams struct {
	Company_Id int
	Agent_Id   int
	Fields     string
	Start_From int
	Limit      int
}
