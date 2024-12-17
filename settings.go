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

type UpdateSysCompanySettingsParams struct {
	Company_Name  string
	Address_1     string
	Address_2     string
	Support_Email string
	Support_Phone string
	Tx            *sql.Tx
}

type UpdateSystenSettingsParams struct {
	Twilio_Account_SID string
	Twilio_Auth_Token  string
	Call_Per_Min       int
	SMS_Per_Page       int
	Credits_Per_Dollar int
	Sendgrid_API       string
	Tx                 *sql.Tx
}

type UpdateAppsInfoParams struct {
	App_ID      int
	Name        string
	Description string
	Price       float64
	Tx          *sql.Tx
}

type UpdatePlanInfoParams struct {
	Plan_ID    int
	Plan_Name  string
	Price      float64
	More_Price float64
	Capacity   int
	Tx         *sql.Tx
}
