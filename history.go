package models

import "database/sql"

type UpdateHistoryParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}
