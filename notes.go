package models

import "database/sql"

type UpdateNoteParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}
