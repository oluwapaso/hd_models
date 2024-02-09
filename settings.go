package models

import "database/sql"

type UpdateDefaultSettParams struct {
	Tx          *sql.Tx
	Feild_Query string
	Query_Value []interface{}
	Where       string
	Must_Update bool
}

type UpdateAccountSettParams struct {
	Tx          *sql.Tx
	Feild_Query string
	Query_Value []interface{}
	Where       string
	Must_Update bool
}

type SystSettParams struct {
	Fields string
}
