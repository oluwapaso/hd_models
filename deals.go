package models

import "database/sql"

type CommonDealFields struct {
	Company_id           int            `json:"company_id"`
	Car_run              string         `json:"car_run"`
	Ship_via             string         `json:"ship_via"`
	Assigned_to          int            `json:"assigned_to"`
	Original_assigned_to sql.NullInt16  `json:"original_assigned_to"`
	Source               sql.NullString `json:"source"`
	Source_type          sql.NullString `json:"source_type"`
	Shipper              sql.NullString `json:"shipper"`
	Referral             sql.NullString `json:"referral"`
	Origin               sql.NullString `json:"origin"`
	Destination          sql.NullString `json:"destination"`
	Vehicles             sql.NullString `json:"vehicles"`
	Payments             sql.NullString `json:"payments"`
	Type                 sql.NullString `json:"type"`
	Original_email       sql.NullString `json:"original_email"`
}

type OrderFields struct {
	Order_id      int            `json:"order_id"`
	Order_uniq_id string         `json:"order_uniq_id"`
	Mode          sql.NullString `json:"mode"`
}

type QuoteFields struct {
	Quote_id      int            `json:"quote_id"`
	Quote_uniq_id int            `json:"quotet_uniq_id"`
	Quote_token   sql.NullString `json:"quote_token"`
}

type Shipper struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone,omitempty"`
	Company string `json:"company,omitempty"`
}

type CommonDates struct {
	Received             string `json:"received"`
	Quoted               string `json:"quoted"`
	First_available_date string `json:"first_available_date"`
	Imported             string `json:"imported,omitempty"`
	DealsDates
}

type DealsDates struct {
	Next_followup_date     string `json:"next_followup_date,omitempty"`     //Quotes only
	Last_updated           string `json:"last_updated,omitempty"`           //Quotes only
	Last_updated_timestamp string `json:"last_updated_timestamp,omitempty"` //Quotes only
	Last_bomb              string `json:"last_bomb,omitempty"`              //Quotes only
	Created                string `json:"created,omitempty"`                //Orders only
}

type Origin struct {
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	Full_origin string `json:"full_origin"`
}

type Destination struct {
	City             string `json:"city"`
	State            string `json:"state"`
	Zip              string `json:"zip"`
	Location         string `json:"location"`
	Address          string `json:"address"`
	Full_destination string `json:"full_destination"`
}

type DealPayments struct {
	Total_tariff      string `json:"total_tariff"`
	Total_brokers_fee string `json:"total_brokers_fee"`
	Carrier_pay       string `json:"carrier_pay"`
}

type Vehicles struct {
	Vehicle_id   int    `json:"vehicle_id"`
	Year         int    `json:"year"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Full_vehicle string `json:"full_vehicle"`
}

type History struct {
	History_Id     int            `json:"history_id"`
	Company_Id     string         `json:"company_id"`
	Item_Type      string         `json:"item_type"`
	Item_Id        int            `json:"item_id"`
	Item_Unique_Id sql.NullString `json:"item_unique_id"`
	History_Type   sql.NullString `json:"history_type"`
	Date_Updated   sql.NullString `json:"date_updated"`
	Message        sql.NullString `json:"message"`
	Changes_Made   sql.NullString `json:"changes_made"`
}
type AddHistoryParams struct {
	Company_Id     string `json:"company_id"`
	Item_Id        int    `json:"item_id"`
	Item_Unique_Id string `json:"item_unique_id"`
	Item_Type      string `json:"item_type"`
	History_Type   string `json:"history_type"`
	Date_Updated   string `json:"date_updated"`
	Updated_By     string `json:"updated_by"`
	Message        string `json:"message"`
}

type Notes struct {
	Note_Id    int            `json:"note_id"`
	Company_Id string         `json:"company_id"`
	Item_Type  string         `json:"item_type"`
	Item_Id    int            `json:"item_id"`
	Note_Type  string         `json:"note_type"`
	Notes      sql.NullString `json:"notes"`
	Date       string         `json:"date"`
	Agent_Id   int            `json:"agent_id"`
	Added_By   string         `json:"added_by"`
}

type AddNotesParams struct {
	Company_Id     string `json:"company_id"`
	Item_Id        int    `json:"item_id"`
	Item_Unique_Id string `json:"item_unique_id"`
	Item_Type      string `json:"item_type"`
	Date           string `json:"date"`
	Agent_Id       int    `json:"agent_id"`
	Note_Type      string `json:"note_type"`
	Notes          string `json:"note"`
	Added_By       string `json:"added_by"`
}

type Payments struct {
	Payment_Id         int            `json:"payment_id"`
	Refrence_No        string         `json:"refrence_no"`
	Order_Id           int            `json:"order_id"`
	Company_Id         string         `json:"company_id"`
	Date_Received      string         `json:"date_received"`
	Payment_From_To    string         `json:"payment_from_to"`
	Amount             string         `json:"amount"`
	Deposit_Method     sql.NullString `json:"deposit_method"`
	Authorization_Code sql.NullString `json:"authorization_code"`
	Transaction_Id     sql.NullString `json:"transaction_id"`
	Notes              sql.NullString `json:"notes"`
	Entered_By         sql.NullString `json:"entered_by"`
}

type AddPaymentsParams struct {
	Refrence_No     string
	Order_Id        int
	Order_Unique_Id string
	Company_Id      string
	Date_Received   string
	Payment_From_To string
	Amount          float64
	Deposit_Method  string
	Transaction_Id  string
	Notes           string
	Entered_By      string
}

type ItemTypeCountsParams struct {
	Company_Id string
	Item_Id    int
	Item_Type  string
	Field      string
	Value      int
}

type DealVehicles struct {
	Vehicle_Id      int         `json:"vehicle_id"`
	Vin             string      `json:"vin"`
	Make            string      `json:"make"`
	Year            interface{} `json:"year"`
	Color           string      `json:"color"`
	Model           string      `json:"model"`
	Type            string      `json:"type"`
	State           string      `json:"state"`
	Tariff          string      `json:"tariff"`
	Running         string      `json:"running"`
	Ship_Via        string      `json:"ship_via"`
	Lot_Number      string      `json:"lot_number"`
	Brokers_Fee     string      `json:"brokers_fee"`
	Full_Vehicle    string      `json:"full_vehicle"`
	Plate_Number    string      `json:"plate_number"`
	Total_Tariff    string      `json:"total_tariff"`
	Additional_Info string      `json:"additional_info"`
}
