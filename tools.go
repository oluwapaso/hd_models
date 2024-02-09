package models

import (
	"database/sql"
	"sync"
)

type CitySearchParams struct {
	Account_Id string
	API_Key    string
	Keyword    string
}

type CityExtended struct {
	City      string `json:"city"`
	State     string `json:"state_code,omitempty"`
	Zip       string `json:"zip"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	County    string `json:"county,omitempty"`
}

type VehicleSearchResp struct {
	Cronetic_Id string `json:"cronetic_id,omitempty"`
	Make        string `json:"make,omitempty"`
	Model       string `json:"model,omitempty"`
}

type PageDataCountParams struct {
	Company_Id   int
	Agent_Id     int
	View_User    string
	View_User_Id int
	Filters      map[string]interface{}
	Page         string
}

type BasicSearchParams struct {
	Company_Id    int
	View_Agent_Id string
	View_User     string
	Keyword       string
	Item_Search   string
	Wg            *sync.WaitGroup
	Ch            chan interface{}
}

type BasicSearchResponse struct {
	LeadsResponse    interface{} `json:"leads_list"`
	QuotesResponse   interface{} `json:"quotes_list"`
	OrdersResponse   interface{} `json:"orders_list"`
	ContactsResponse interface{} `json:"contacts_list"`
}

type AdvancedSearchParams struct {
	Company_Id   int
	Fields       string
	Account_Type string
	Search_Query string
	Params       []interface{}
	StartFrom    int
	Limit        int
}

type UnsubParams struct {
	Email      string
	Company_Id int
	Date       string
	Tx         *sql.Tx
}
