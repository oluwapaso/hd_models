package models

import "database/sql"

type ImportedDealFiles struct {
	Import_id        int            `json:"import_id"`
	Company_id       string         `json:"company_id"`
	File_location    string         `json:"file_location,omitempty"`
	Data_headers     string         `json:"data_headers,omitempty"`
	Data_rows        string         `json:"data_rows,omitempty"`
	Data_type        string         `json:"data_type,omitempty"`
	Uploaded_by      string         `json:"uploaded_by,omitempty"`
	Default_status   sql.NullString `json:"default_status,omitempty"`
	Default_agent    sql.NullInt32  `json:"default_agent,omitempty"`
	Rejected_leads   int            `json:"rejected_leads,omitempty"`
	Total_file_data  int            `json:"total_file_data,omitempty"`
	Imported_data    int            `json:"imported_data,omitempty"`
	Can_start_import string         `json:"can_start_import,omitempty"`
	Imported         string         `json:"imported,omitempty"`
	Date_added       string         `json:"date_added,omitempty"`
}

type ImportResp struct {
	RepoRes RepoResponse
	Data    ImportedDealFiles
}

type ImportDealHeader struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Skip         string `json:"skip"`
	Matched      string `json:"matched"`
	Matched_with string `json:"matched_with"`
}

type ImportedStatus struct {
	Map_id                string `json:"map_id"`
	Company_id            string `json:"company_id"`
	Mapped_quote_status   string `json:"mapped_quote_status"`
	Mapped_order_status   string `json:"mapped_order_status"`
	Mapped_contact_status string `json:"mapped_contact_status"`
}

type ImportedQuoteStatus struct {
	Active_quote       string `json:"active_quote"`
	Contacted_quote    string `json:"contacted_quote"`
	Responsive_quote   string `json:"responsive_quote"`
	Unresponsive_quote string `json:"unresponsive_quote"`
	Onhold_quote       string `json:"onhold_quote"`
}

type CommonImportHeaders struct {
	Status         string `json:"status"`
	Name           string `json:"name"`
	Company        string `json:"company"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Vehicles       string `json:"vehicles"`
	Origin         string `json:"origin"`
	Destination    string `json:"destination"`
	Ship_Date      string `json:"ship_date"`
	Transport_Type string `json:"transport_type"`
	Total_Tariff   string `json:"total_tariff"`
	Carrier_Pay    string `json:"carrier_pay"`
	Broker_Fee     string `json:"broker_fee"`
	Assigned_To    string `json:"assigned_to"`
	Source         string `json:"source"`
}

type ImportedQuoteHeaders struct {
	CommonImportHeaders
	Quote_Number  string `json:"quote_number"`
	Quote_Created string `json:"quote_created"`
	Followup_Date string `json:"followup_date"`
}

type ImportedOrderStatus struct {
	Active_order string `json:"active_order"`
	Onhold_order string `json:"onhold_order"`
}

type ImportedContactStatus struct {
	Active_contact   string `json:"active_contact"`
	Inactive_contact string `json:"inactive_contact"`
}

type ImportedOrderHeaders struct {
	CommonImportHeaders
	Order_Number  string `json:"order_number"`
	Order_Created string `json:"order_created"`
}

type MappedImportStatus struct {
	System_status string `json:"system_status"`
	Mapped_status string `json:"mapped_status"`
}

type ImportedStatusResp struct {
	RepoRes RepoResponse
	Data    ImportedStatus
}
