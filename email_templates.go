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
	Email_Unique_Id int         `json:"email_unique_id"`
	Email_Id        int         `json:"email_id"`
	Company_Id      string      `json:"company_id,omitempty"`
	Email_Type      string      `json:"email_type,omitempty"`
	Name            string      `json:"name"`
	Used_For        string      `json:"used_for"`
	Stats_Data      interface{} `json:"stats_data"`
	For_Follow_Up   string      `json:"for_follow_up"`
	Html_Contents   string      `json:"html_contents,omitempty"`
	Subject         string      `json:"subject,omitempty"`
	Description     string      `json:"description,omitempty"`
	Attachment      interface{} `json:"attachment,omitempty"`
	CC_Addresses    string      `json:"cc_addresses,omitempty"`
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
	Used_For   string
}

type SingleEmlTempParams struct {
	Type          string
	Company_Id    int
	Template_Id   int
	Template_Name string
	Used_For      string
	Temp_Type     string
	Fields        string
	Order_By      string
	Limit         string
}

type DefaultEmlTempParams struct {
	Template_Name string
	Fields        string
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
	Tx            *sql.Tx
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
	Charged             string `json:"charged"`
	Unique_Opens        string `json:"unique_opens"`
	Last_Click          string `json:"last_click"`
	Last_Open           string `json:"last_open"`
	Last_Unique_Click   string `json:"last_unique_click"`
	Last_Unique_Open    string `json:"last_unique_open"`
	CC_Addresses        string `json:"cc_addresses"`
	BCC_Addresses       string `json:"bcc_addresses"`
	Tracking_Id         int    `json:"tracking_id"`
	Template_Id         int    `json:"template_id"`
	Automation_UID      string `json:"automation_uid"`
	Automation_Step_UID string `json:"automation_step_uid"`
}

type UpdateEmailTemplateParams struct {
	Company_Id       int
	Template_Id      int
	Temp_Name        string
	Use_For_Followup string
	Used_For         string
	Email_Body       string
	Subject          string
	Description      string
	Attachments      string
	CC_Addresses     string
	Tx               *sql.Tx
}

type AddEmailTemplateParams struct {
	Company_Id    int
	Next_Email_Id int
	Email_Name    string
	Used_For      string
	Tx            *sql.Tx
}

type DeleteEmailTemplateParams struct {
	Company_Id  int
	Template_Id int
	Tx          *sql.Tx
}

type LoadBatchEmailParams struct {
	Company_Id int
	Batch_UID  string
	Fields     string
}

type CountTempSentMsgsParams struct {
	Message_Type string
	Company_Id   int
	Template_Id  int
	Tx           *sql.Tx
}

type UpdateTempSentCountsParams struct {
	Company_Id  int
	Template_Id int
	Total_Sent  int
	Tx          *sql.Tx
}
