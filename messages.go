package models

import "database/sql"

type Messages struct {
	Message_id        int            `json:"message_id"`
	Agent_id          string         `json:"agent_id"`
	Company_id        string         `json:"company_id"`
	Subject           sql.NullString `json:"subject"`
	Preview           sql.NullString `json:"preview"`
	Contact_email     sql.NullString `json:"contact_email"`
	Contact_phone     sql.NullString `json:"contact_phone"`
	Contact_name      sql.NullString `json:"contact_name"`
	Contact_type      sql.NullString `json:"contact_type"`
	Has_thread        string         `json:"has_thread"`
	Status            string         `json:"status"`
	New_messages      int            `json:"new_messages"`
	Last_message_type string         `json:"last_message_type"`
	Last_unique_id    sql.NullString `json:"last_unique_id"`
	Last_message_date sql.NullString `json:"last_message_date"`
}

type MessageThreads struct {
	Thread_id            int            `json:"thread_id,omitempty"`
	Message_id           string         `json:"message_id,omitempty"`
	Unique_id            string         `json:"unique_id,omitempty"`
	Message_type         string         `json:"message_type,omitempty"`
	Agent_id             string         `json:"agent_id,omitempty"`
	Deal_type            string         `json:"deal_type,omitempty"`
	Deal_id              sql.NullString `json:"deal_id,omitempty"`
	Company_id           string         `json:"company_id,omitempty"`
	Contact_name         sql.NullString `json:"contact_name,omitempty"`
	Contact_type         string         `json:"contact_type,omitempty"`
	Forwarding_email     sql.NullString `json:"forwarding_email,omitempty"`
	Send_from            string         `json:"send_from,omitempty"`
	Send_to              string         `json:"send_to,omitempty"`
	Subject              sql.NullString `json:"subject,omitempty"`
	Message_body         sql.NullString `json:"message_body,omitempty"`
	Date                 string         `json:"date,omitempty"`
	Last_send_attempt    sql.NullString `json:"last_send_attempt,omitempty"`
	Sent_by              string         `json:"sent_by,omitempty"`
	Status               string         `json:"status,omitempty"`
	Outgoing_mail_status string         `json:"outgoing_mail_status,omitempty"`
	Outgoing_mail_info   sql.NullString `json:"outgoing_mail_info,omitempty"`
	Batch_UID            string         `json:"batch_uid,omitempty"`
}

type EmailQueues struct {
	Queue_id     int            `json:"queue_id"`
	Type         string         `json:"type"`
	Company_id   string         `json:"company_id"`
	Company_name sql.NullString `json:"company_name"`
	Agent_id     sql.NullString `json:"agent_id"`
	Agent_name   sql.NullString `json:"agent_name"`
	Subject      string         `json:"subject"`
	Email_body   string         `json:"body"`
	From_email   string         `json:"from_email"`
	To_email     string         `json:"to_email"`
	To_name      sql.NullString `json:"to_name"`
	Notify       string         `json:"notify"`
	Status       string         `json:"status"`
	More_info    sql.NullString `json:"more_info"`
	Date_added   string         `json:"date_added"`
}

type MessageRepoResponse struct {
	Resp RepoResponse
	Data Messages
}

type MessageThreadRepoResponse struct {
	Resp RepoResponse
	Data []MessageThreads
}

type AddEmailToMessages struct {
	Company_Id               string `json:"company_id"`
	To_Email                 string `json:"to_email"`
	Agent_Id                 int    `json:"agent_id"`
	Agent_Email              string `json:"agent_email"`
	Subject                  string `json:"subject"`
	Body                     string `json:"body"`
	Deal_Type                string `json:"deal_type"` //None
	Deal_Id                  int    `json:"deal_id"`
	Deal_Unique_Id           string `json:"deal_unique_id"`
	Tracking_Id              int    `json:"tracking_id"`
	Template_Id              int    `json:"template_id"`
	Composer                 string `json:"composer"`        //none
	Composer_Msg_Id          string `json:"composer_msg_id"` //Used as alternate contact
	Contact_Name             string `json:"contact_name"`
	Contact_Phone            string `json:"contact_phone"`
	Contact_Type             string `json:"contact_type"` //NormalShipper
	CC_Addresses             string `json:"cc_addresses"`
	BBC_Addresses            string `json:"bcc_addresses"`
	Attachments              string `json:"attachments"`
	Automation_UID           string `json:"automation_uid"`
	Automation_Step_UID      string `json:"automation_step_uid"`
	Agent_Forwarding_Address string `json:"agent_forwarding_address"`
	Status                   string `json:"status"`  //Delivered
	Sent_by                  string `json:"sent_by"` //Agent
	Batch_UID                string `json:"batch_uid"`
}

type AddSMSToMessages struct {
	Company_Id               string `json:"company_id"`
	From_Number              string `json:"from_number"`
	To_Number                string `json:"to_number"`
	Agent_Id                 int    `json:"agent_id"`
	Agent_Email              string `json:"agent_email"`
	Body                     string `json:"body"`
	Deal_Type                string `json:"deal_type"` //None
	Deal_Id                  int    `json:"deal_id"`
	Deal_Unique_Id           string `json:"deal_unique_id"`
	Template_Id              int    `json:"template_id"`
	Composer                 string `json:"composer"` //none
	Composer_Msg_Id          string `json:"composer_msg_id"`
	Contact_Name             string `json:"contact_name"`
	Contact_Email            string `json:"contact_email"` //Used as alternate contact
	Contact_Type             string `json:"contact_type"`  //NormalShipper
	Automation_UID           string `json:"automation_uid"`
	Automation_Step_UID      string `json:"automation_step_uid"`
	Agent_Forwarding_Address string `json:"agent_forwarding_address"`
	Status                   string `json:"status"`  //Delivered
	Sent_by                  string `json:"sent_by"` //Agent
}

type AddClientMsgInfoParams struct {
	Company_Id           string `json:"company_id"`
	Agent_Id             int    `json:"agent_id"`
	Contact_Email        string `json:"contact_email"`
	Contact_Phone_1      string `json:"contact_phone_1"`
	Contact_Phone_2      string `json:"contact_phone_2"`
	Contact_Name         string `json:"contact_name"`
	Contact_Account_Type string `json:"contact_accn_type"`
}

type GetOrAddClientMsgInfoParams struct {
	Company_Id           string `json:"company_id"`
	Agent_Id             int    `json:"agent_id"`
	Contact_Info         string `json:"contact_info"`
	Alt_Contact_Info     string `json:"alt_contact_info"`
	Contact_Type         string `json:"contact_type"`
	Contact_Name         string `json:"contact_name"`
	Contact_Account_Type string `json:"contact_accn_type"`
}

type AddEmailQueueParams struct {
	Type         string `json:"type"`
	Company_Id   int    `json:"company_id"`
	Company_Name string `json:"company_name"`
	Agent_Id     int    `json:"agent_id"`
	Agent_Name   string `json:"agent_name"`
	Subject      string `json:"subject"`
	Body         string `json:"body"`
	From_Email   string `json:"from_email"`
	To_Email     string `json:"to_email"`
	To_Name      string `json:"to_name"`
	Notify       string `json:"notify"`
	More_Info    string `json:"more_info"`
	Expires_On   string `json:"expires_on"`
}

type AddSMSQueueParams struct {
	Type        string `json:"type"`
	Company_Id  int    `json:"company_id"`
	Agent_Id    int    `json:"agent_id"`
	Agent_Name  string `json:"agent_name"`
	Body        string `json:"body"`
	From_Number string `json:"from_number"`
	To_Number   string `json:"to_number"`
	More_Info   string `json:"more_info"`
}

type UpdateMsgThreadParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}

type ListsDealThreadsParams struct {
	Company_Id int
	Fields     string
	Item_Type  string
	Item_Id    int
	Limit      int
}

type GetEmailReportsParams struct {
	Company_Id   int
	View_Mode    string
	View_User_Id int
	Item_Type    string
	Batch_UID    string
	Fields       string
	StartFrom    int
	Limit        int
}

type ListMessagesParams struct {
	Company_Id  int
	Agent_Id    int
	Filter_Type string
	Fields      string
	Search_For  string
	Start_From  int
	Limit       int
}

type UpdateMsgStatusParams struct {
	Company_Id  int
	Agent_Id    int
	Status      string
	Message_IDs string
	Tx          *sql.Tx
}

type ListsMsgThreadsParams struct {
	Company_Id int
	Agent_Id   int
	Fields     string
	Message_Id int
	Start_From string
	Limit      int
}

type AddBatchEmailParams struct {
	Company_Id int
	Agent_Id   int
	Batch_UID  string
	Subject    string
	Mail_Info  string
	Date       string
	Tx         *sql.Tx
}

type LoadSingleThreadMsgParams struct {
	Company_Id  int
	Load_Type   string
	Tracking_Id int
	Fields      string
}

type UpdateThreadOpensParams struct {
	Update_Type           string
	Company_Id            int
	Tracking_Id           int
	Thread_Id             int
	New_Opens             int
	First_Open_Date       string
	Last_Open_Date        string
	Last_Unique_Open_Date string
	Tx                    *sql.Tx
}

type UpdateBatchOpensParams struct {
	Update_Type           string
	Company_Id            int
	Batch_Id              int
	New_Opens             int
	Last_Open_Date        string
	Last_Unique_Open_Date string
	Tx                    *sql.Tx
}

type UpdateTempOpensParams struct {
	Update_Type           string
	Company_Id            int
	Temp_Id               int
	New_Opens             int
	Last_Open_Date        string
	Last_Unique_Open_Date string
	Tx                    *sql.Tx
}

type UpdateThreadClicksParams struct {
	Update_Type            string
	Company_Id             int
	Tracking_Id            int
	Thread_Id              int
	New_Clicks             int
	Last_Click_Date        string
	Last_Unique_Click_Date string
	Tx                     *sql.Tx
}

type UpdateBatchClicksParams struct {
	Update_Type            string
	Company_Id             int
	Batch_Id               int
	New_Clicks             int
	Last_Click_Date        string
	Last_Unique_CLick_Date string
	Tx                     *sql.Tx
}

type UpdateTempClicksParams struct {
	Update_Type            string
	Company_Id             int
	Temp_Id                int
	New_Cicks              int
	Last_Click_Date        string
	Last_Unique_Click_Date string
	Tx                     *sql.Tx
}

type AddForwardingEmailParams struct {
	Company_Id int
	Agent_Id   int
	Forwarder  string
	Tx         *sql.Tx
}

type CheckMsgParams struct {
	Check_By    string
	Company_Id  int
	Agent_Id    int
	From_Email  string
	From_Number string
	Fields      string
}

type CheckExistingThrdParams struct {
	Check_Type         string
	Company_Id         int
	Agent_Id           int
	Unique_Id          string
	Forwarding_Address string
}

type AddMsgToThreadParams struct {
	Message_Type       string
	Message_Id         string
	Unique_Id          string
	Company_Id         int
	Agent_Id           int
	Contact_Name       string
	Forwarding_Address string
	From_Email         string
	To_Email           string
	From_Number        string
	To_Number          string
	Subject            string
	Body               string
	Date               string
	Item_Id            int
	Item_UID           string
	Item_Type          string
	Contact_Type       string
	Tx                 *sql.Tx
}

type UpdateMsgParams struct {
	Update_Type        string
	Message_Id         string
	Unique_Id          string
	Contact_Name       string
	Forwarding_Address string
	Subject            string
	Preview            string
	New_Msg            int
	Date               string
	Tx                 *sql.Tx
}
