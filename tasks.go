package models

import "database/sql"

type UpdateTaskParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}
