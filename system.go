package models

import "database/sql"

type SystemReports struct {
	Report_Id   int    `json:"report_id"`
	Company_Id  int    `json:"company_id"`
	Agent_Id    int    `json:"agent_id"`
	Report_Type string `json:"report_type"`
	Page        string `json:"page"`
	Message     string `json:"message"`
	Date        string `json:"date"`
}

type CheckPageTutNameParams struct {
	Name string
}

type AddPageTutParams struct {
	Page_Name  string
	Video_Link string
	Tx         *sql.Tx
}

type ListPageTutorialsParams struct {
	Type      string
	Fields    string
	Page_Name string
	StartFrom int
	Limit     int
}

type DeletePageTutorialsParams struct {
	Page_ID int
	Tx      *sql.Tx
}

type DeleteSysUpdateParams struct {
	Update_ID int
	Tx        *sql.Tx
}

type AddSysUpdateParams struct {
	Update_Type string
	Update_Body string
	Tx          *sql.Tx
}

type AddSysAlertParams struct {
	Alert_Type string
	Alert_Body string
	Tx         *sql.Tx
}

type DeleteSysAlertParams struct {
	Alert_ID int
	Tx       *sql.Tx
}
