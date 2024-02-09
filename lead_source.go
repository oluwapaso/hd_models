package models

import "database/sql"

type LeadSource struct {
	Source_id           int            `json:"source_id"`
	Company_id          string         `json:"company_id"`
	Source_email        string         `json:"source_email"`
	API_access_key      string         `json:"api_access_key"`
	Status              string         `json:"status"`
	Distributed         string         `json:"distributed"`
	Company_info        string         `json:"company_info"`
	Company_name        sql.NullString `json:"company_name"` //E.g json_extract(company_info, '$.name') AS company_name
	Auto_quote          string         `json:"auto_quote"`
	Autoqouote_formular string         `json:"autoqouote_formular"`
	Use_for_extnl_form  string         `json:"use_for_extnl_form"`
	Deflt_price_checker string         `json:"deflt_price_checker"`
	Thank_you_page      string         `json:"thank_you_page"`
	Settings            string         `json:"settings"`
	Import_mapping      sql.NullString `json:"import_mapping"`
	Got_Mails           string         `json:"got_mails"`
}

type ImportedLeadSourceResp struct {
	RepoRes RepoResponse
	Data    []LeadSource
}

type MappedLeadSource struct {
	Source_id     int    `json:"source_id"`
	Company_name  string `json:"company_name"`
	Mapped_import string `json:"mapped_import"`
}

type DeleteEmailResponse struct {
	Success            bool   `json:"success"`
	Error_Message      string `json:"error_message"`
	No_Of_Inbox_Msgs   int    `json:"no_of_inbox_msgs"`
	No_Of_Deleted_Msgs int    `json:"no_of_deleted_msgs"`
	Source_Email       string `json:"source_email"`
}

type LeadSourceListsParams struct {
	Type       string
	Company_Id int
	Fields     string
	Order_By   string
	StartFrom  int
	Limit      int
}

type LoadSingleLeadSourceParams struct {
	Company_Id     int
	Fields         string
	Lead_Source_Id int
}

type CountLeadSourceDealsParams struct {
	Source_Company_Name string
	LoadReportsParams
}

type LeadSrcDispDataParams struct {
	Source_Company_Name string
	LoadReportsParams
}

type LeadSourceDealsResponse struct {
	Leads            int
	Quotes           int
	Orders           int
	Leads_To_Quotes  int
	Leads_To_Orders  int
	Quotes_To_Orders int
	Created_Quotes   int
	Created_Orders   int
}

type LeadSrcDispDataResponse struct {
	Total_Disp_Orders int
	TotalBrokersFee   interface{}
	TotalTariff       interface{}
	TotalCarrierPay   interface{}
}

type DealsReportData struct {
	Current_Leads          int     `json:"total_current_leads"`
	Current_Quotes         int     `json:"total_current_quotes"`
	Current_Orders         int     `json:"total_current_orders"`
	Total_Leads            int     `json:"total_leads"`
	Total_Quotes           int     `json:"total_quotes"`
	Total_Orders           int     `json:"total_orders"`
	Dispatched_Orders      int     `json:"dispatched_orders"`
	Total_Tariff           float64 `json:"total_tariff"`
	Conversion_Rate        float64 `json:"conversion_rate"`
	Leads_Conversion_Rate  float64 `json:"leads_conversion_rate"`
	Quotes_Conversion_Rate float64 `json:"quotes_conversion_rate"`
	Total_Carrier_Pay      float64 `json:"total_carrier_pay"`
	Total_Brokers_Fee      float64 `json:"total_brokers_fee"`
	Gross_Profit           float64 `json:"gross_profit"`
	Profit_Margin          float64 `json:"profit_margin"`
	Avg_Price_Per_Lead     float64 `json:"avg_profit_per_lead"`
}

type LeadSourceReportData struct {
	LeadSource         string  `json:"lead_source"`
	Lead_Source_Status string  `json:"status"`
	Total_Leads_Amnt   float64 `json:"total_leads_amount"`
	DealsReportData
}

type UnsetExtrnlSRCParams struct {
	Company_Id int
	Tx         *sql.Tx
}

type UpdateLeadSRCParams struct {
	Tx          *sql.Tx
	Feild_Query string
	Query_Value []interface{}
	Where       string
	Must_Update bool
}

type MappedSrcParams struct {
	Company_Id int
	Fields     string
}

type UpdateImportMappedStatusParams struct {
	Company_Id     int
	Quote_Status   string
	Order_Status   string
	Contact_Status string
	Must_Update    bool
	Tx             *sql.Tx
}

type MappedAgentParams struct {
	Company_Id int
	Value      string
	Agent_Id   int
}

type MultiMappedSrcParams struct {
	Company_Id int
	Value      string
	Source_Id  int
}

type GetSrcDetailsParams struct {
	Company_Id  int
	Fields      string
	Lead_Medium string
	Query_By    string
}

type AssignLeadToParams struct {
	Company_Id    int
	Source_Id     int
	Shipper_Email string
	Tx            *sql.Tx
}

type AssignLeadToResponse struct {
	Lead_Multiples         int
	Route_Leads            string
	Leads_Assigned         int
	Agent_Id               int
	Name                   string
	Phone                  string
	Email                  string
	Username               string
	Notifications_Settings string
	Mailer_Settings        string
	Webmail_Settings       string
}

type UpdateAss_RR_InfoParams struct {
	Assign_Next_Lead string
	Leads_Assigned   int
	Company_Id       int
	Source_Id        int
	Agent_Id         int
	Tx               *sql.Tx
}

type CheckLeadSourceParam struct {
	Company_Id   int
	Check_By     string
	Source_Id    int
	Source_Email string
}

type CheckAgentAndSrcParam struct {
	Company_Id int
	Source_Id  int
	Agent_Id   int
}

type CheckDistributedParam struct {
	Company_Id int
	Source_Id  int
	Agent_Id   int
}

type CheckSRCDistributedParam struct {
	Company_Id int
	Source_Id  int
	Tx         *sql.Tx
}

type UpSrcSettParams struct {
	Company_Id     int
	Source_Id      int
	Agent_Id       int
	Lead_Multiples int
	Route_To       string
	Tx             *sql.Tx
}

type AddSrcSettParams struct {
	Company_Id     int
	Source_Id      int
	Agent_Id       int
	Lead_Multiples int
	Route_To       string
	Tx             *sql.Tx
}

type UpdateDistributedParam struct {
	Company_Id  int
	Source_Id   int
	Distributed string
	Tx          *sql.Tx
}

type ResetPCParams struct {
	Company_Id int
	Tx         *sql.Tx
}

type AddLeadSrcParams struct {
	Company_Id          int
	Source_Email        string
	Email_Password      string
	API_Access_Key      string
	Price_Per_Lead      string
	Company_Info        string
	Auto_Quote_Formular string
	Settings            string
	Tx                  *sql.Tx
}

type SetGotMailsParams struct {
	Source_Email string
	Tx           *sql.Tx
}

type DefaultPriceCheckerAqFormularParams struct {
	Company_Id int
}
