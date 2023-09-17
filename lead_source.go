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
