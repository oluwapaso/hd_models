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
	Quote_token          sql.NullString `json:"quote_token"`
}

type OrderFields struct {
	Order_id      int            `json:"order_id"`
	Order_uniq_id int            `json:"order_uniq_id"`
	Mode          sql.NullString `json:"mode"`
}

type QuoteFields struct {
	Quote_id      int `json:"quote_id"`
	Quote_uniq_id int `json:"quotet_uniq_id"`
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

type Payments struct {
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
