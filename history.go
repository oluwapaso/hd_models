package models

import "database/sql"

type UpdateHistoryParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
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

type ListsDealHistoryParams struct {
	Company_Id int
	Fields     string
	Item_Type  string
	Item_Id    string
}
