package models

type DB_Query_Response struct {
	Total_Records  int    `json:"total_records"`
	Total_Returned int    `json:"total_returned"`
	Has_More       string `json:"has_more"`
}
