package models

import "database/sql"

type UpdateErrorParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}

type ListsErrorsParams struct {
	Company_Id    int
	Fields        string
	StartFrom     int
	Limit         int
	Initiator     string
	Item_Id       string
	View_User     string
	View_Agent_Id string
}

type ListsDealErrorsParams struct {
	Company_Id int
	Fields     string
	Item_Type  string
	Item_Id    string
}
