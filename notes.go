package models

import "database/sql"

type UpdateNoteParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
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
	Company_Id         string `json:"company_id"`
	Item_Id            int    `json:"item_id"`
	Item_Unique_Id     string `json:"item_unique_id"`
	Item_Type          string `json:"item_type"`
	Date               string `json:"date"`
	Agent_Id           int    `json:"agent_id"`
	Note_Type          string `json:"note_type"`
	Notes              string `json:"note"`
	Added_By           string `json:"added_by"`
	Update_Deals_Count string `json:"update_deals_count"`
}
