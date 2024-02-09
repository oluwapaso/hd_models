package models

import "database/sql"

type ListSMSTemplatesParams struct {
	Type       string
	Company_Id int
	Type_Value string
	Fields     string
	StartFrom  int
	Limit      int
	Used_For   string
}

type SingleSMSTempParams struct {
	Type          string
	Company_Id    int
	Template_Id   int
	Template_Name string
	Fields        string
	Order_By      string
	Limit         string
}

type SMSTempLists struct {
	SMS_Unique_Id int    `json:"sms_unique_id"`
	SMS_Id        int    `json:"sms_id"`
	Company_Id    string `json:"company_id"`
	Name          string `json:"name"`
	Used_For      string `json:"used_for,omitempty"`
	SMS_Body      string `json:"sms_body,omitempty"`
	Description   string `json:"description,omitempty"`
}

type SMS_Temp_Lists_Resp struct {
	DB_Query_Response
	SMS_Templates interface{} `json:"sms_templates"`
}

type AddSMSTemplateParams struct {
	Company_Id   int
	Next_Temp_Id int
	SMS_Name     string
	Used_For     string
	SMS_Body     string
	Tx           *sql.Tx
}

type UpdateSMSTemplateParams struct {
	Company_Id  int
	Template_Id int
	Temp_Name   string
	SMS_Body    string
	Description string
	Tx          *sql.Tx
}

type DeleteSMSTemplateParams struct {
	Company_Id  int
	Template_Id int
	Tx          *sql.Tx
}
