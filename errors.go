package models

import "database/sql"

type UpdateErrorParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}
