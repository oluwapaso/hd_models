package models

type MailerEvent struct {
	Server        string `json:"server"`
	Thread_ID     int    `json:"thread_id"`
	From_Email    string `json:"from_email"`
	From_Name     string `json:"from_name"`
	To_Email      string `json:"to_email"`
	To_Name       string `json:"to_name"`
	Subject       string `json:"subject"`
	Email_Body    string `json:"email_body"`
	CC_Addresses  string `json:"cc_addresses"`
	SMTP_Host     string `json:"smtp_host,omitempty"`
	SMTP_Port     int    `json:"smtp_port,omitempty"`
	SMTP_Username string `json:"smtp_username,omitempty"`
	SMTP_Password string `json:"smtp_password,omitempty"`
	SendGrid_Key  string `json:"sendgrid_key,omitempty"`
	Company_Id    string `json:"company_id"`
	Agent_Id      int    `json:"agent_id"`
	Deal_Type     string `json:"deal_type"`
	Deal_Id       int    `json:"deal_id"`
}

type SMSEvent struct {
	Company_Id         string `json:"company_id"`
	From               string `json:"from"`
	To                 string `json:"to"`
	Message_Body       string `json:"message_body"`
	Thread_ID          int    `json:"thread_id"`
	Twilio_Account_SID string `json:"twilio_account_sid"`
	Twilio_Auth_Token  string `json:"twilio_auth_token"`
	Twilio_Credits     int    `json:"tiwilo_credits"`
	Agent_Id           int    `json:"agent_id"`
	Deal_Type          string `json:"deal_type"`
	Deal_Id            int    `json:"deal_id"`
}

type LambdaResponse struct {
	StatusCode   string `json:"status_code"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
}

type LambdaError struct {
	Message   string `json:"message"`
	Thread_Id int    `json:"thread_id"`
}

type LambdaSendEmailResp struct {
	Message    string `json:"message"`
	Thread_ID  int    `json:"thread_id"`
	Company_Id string `json:"company_id,omitempty"`
	Agent_Id   int    `json:"agent_id,omitempty"`
	Deal_Type  string `json:"deal_type,omitempty"`
	Deal_Id    int    `json:"deal_id,omitempty"`
}

type LambdaSendSMSResp struct {
	Message        string `json:"message"`
	Thread_ID      int    `json:"thread_id"`
	Num_Of_Page    int    `json:"num_of_page,omitempty"`
	Twilio_Credits int    `json:"twilio_credits,omitempty"`
	Company_Id     string `json:"company_id,omitempty"`
	Agent_Id       int    `json:"agent_id,omitempty"`
	Deal_Type      string `json:"deal_type,omitempty"`
	Deal_Id        int    `json:"deal_id,omitempty"`
}

type LambdaInvokeOutput struct {
	Log_Message string `json:"log_message"`
	Thread_ID   int    `json:"thread_id"`
	Company_Id  string `json:"company_id,omitempty"`
	Agent_Id    int    `json:"agent_id,omitempty"`
	Deal_Type   string `json:"deal_type,omitempty"`
	Deal_Id     int    `json:"deal_id,omitempty"`
}

type IMAP_Event struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
type LambdaDelOldEmlOutput struct {
	Source_Email       string `json:"source_email"`
	No_Of_Inbox_Msgs   int    `json:"no_of_inbox_msg"`
	No_Of_Deleted_Msgs int    `json:"no_of_deleted_msg"`
}

type Delete_Src_Emails_Error struct {
	Message      string `json:"message"`
	Source_Email string `json:"source_email"`
}
