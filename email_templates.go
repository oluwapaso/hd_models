package models

import (
	"database/sql"
)

type EmailTemplate struct {
	Email_Unique_Id int            `json:"email_unique_id"`
	Email_Id        int            `json:"email_id"`
	Company_Id      string         `json:"company_id"`
	Email_type      string         `json:"email_type"`
	Name            string         `json:"name"`
	Used_For        string         `json:"used_for"`
	For_Follow_Up   string         `json:"for_follow_up"`
	Html_Contents   sql.NullString `json:"html_contents"`
	Subject         sql.NullString `json:"subject"`
	Description     sql.NullString `json:"description"`
	Attachment      sql.NullString `json:"attachment"`
	CC_Addresses    sql.NullString `json:"cc_addresses"`
}

// Used for dynamic fetching. So sql.NullString and likes will not be used here
type EmailTempLists struct {
	Email_Unique_Id int    `json:"email_unique_id"`
	Email_Id        int    `json:"email_id"`
	Company_Id      string `json:"company_id,omitempty"`
	Email_type      string `json:"email_type,omitempty"`
	Name            string `json:"name"`
	Used_For        string `json:"used_for"`
	For_Follow_Up   string `json:"for_follow_up"`
	Html_Contents   string `json:"html_contents,omitempty"`
	Subject         string `json:"subject,omitempty"`
	Description     string `json:"description,omitempty"`
	Attachment      string `json:"attachment,omitempty"`
	CC_Addresses    string `json:"cc_addresses,omitempty"`
}

type EmailTempListsResponse struct {
	Total_Records   int         `json:"total_records"`
	Total_Returned  int         `json:"total_returned"`
	Has_More        string      `json:"has_more"`
	Email_Templates interface{} `json:"email_templates"`
}

type SelEmlTempParams struct {
	Company_Id      string
	Email_Unique_Id int
	Search_By       string
	Fields          string
}

type ListEmailTemplatesParams struct {
	Type       string
	Company_Id int
	Type_Query string
	Type_Value string
	Fields     string
	StartFrom  int
	Limit      int
}

type TempReplaceParams struct {
	Company_Id    string
	Item_Id       int
	Item_Type     string
	Template_Type string
	Template_Id   int
	Subject       string
	Body          string
	Agent_Id      int
	Agent_Name    string
	Agent_Email   string
	Agent_Phone   string
}

type ReplacedEmailTemp struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type TrackerParams struct {
	Company_Id string
	Agent_Id   int
	Item_Id    int
	Item_Type  string
}

type OutgoingMailInfo struct {
	Send_Attempts       string `json:"send_attempts"`
	Clicks              string `json:"clicks"`
	Unique_Clicks       string `json:"unique_clicks"`
	Opens               string `json:"opens"`
	Unique_Opens        string `json:"unique_opens"`
	Last_Unique_Click   string `json:"last_unique_click"`
	Last_Unique_Open    string `json:"last_unique_open"`
	CC_Addresses        string `json:"cc_addresses"`
	BCC_Addresses       string `json:"bcc_addresses"`
	Tracking_Id         int    `json:"tracking_id"`
	Template_Id         int    `json:"template_id"`
	Automation_UID      string `json:"automation_uid"`
	Automation_Step_UID string `json:"automation_step_uid"`
}
