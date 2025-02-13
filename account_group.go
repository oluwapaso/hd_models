package models

import "database/sql"

type LoadAccountGroupsParams struct {
	Load_Type  string
	Company_Id int
	Fields     string
	Order_By   string
	Group_Name string
	Group_Id   int
	Paginated  string
	Start_From int
	Limit      int
}

type AddAccountGroupsParams struct {
	Company_Id         int
	Group_Name         string
	Group_Descriptions string
	Permissions        string
	Tx                 *sql.Tx
}

type UpdateAccountGroupsParams struct {
	Company_Id         int
	Group_Id           int
	Group_Name         string
	Group_Descriptions string
	Permissions        string
	Tx                 *sql.Tx
}

type DeleteAccountGroupsParams struct {
	Company_Id int
	Group_Id   int
	Tx         *sql.Tx
}
