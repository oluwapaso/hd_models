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
