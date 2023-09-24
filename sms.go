package models

type ListSMSTemplatesParams struct {
	Type       string
	Company_Id int
	Type_Value string
	Fields     string
	StartFrom  int
	Limit      int
}

type SingleSMSTempParams struct {
	Type        string
	Company_Id  int
	Template_Id int
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
