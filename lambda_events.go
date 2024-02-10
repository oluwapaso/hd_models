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
	BCC_Addresses string `json:"bcc_addresses"`
	Attachments   string `json:"attachments"`
	Files         string `json:"files"`
	SMTP_Host     string `json:"smtp_host,omitempty"`
	SMTP_Port     int    `json:"smtp_port,omitempty"`
	SMTP_Username string `json:"smtp_username,omitempty"`
	SMTP_Password string `json:"smtp_password,omitempty"`
	SendGrid_Key  string `json:"sendgrid_key,omitempty"`
	Company_Id    string `json:"company_id"`
	Agent_Id      int    `json:"agent_id"`
	Deal_Type     string `json:"deal_type"`
	Deal_Id       int    `json:"deal_id"`
	Batch_UID     string `json:"batch_uid"`
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

type Lambda_API_Response struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
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
	Batch_UID  string `json:"batch_uid,omitempty"`
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
	Batch_UID   string `json:"batch_uid"`
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

type Page_PDF_Event struct {
	Page_Link string `json:"page_link"`
}

type CD_Lambda_Event struct {
	Company_Id        int    `json:"company_id"`
	Agent_Id          int    `json:"agent_id"`
	Resource          string `json:"resource"`
	CD_Username       string `json:"cd_username"`
	CD_Password       string `json:"cd_password"`
	CD_Search_Cookies string `json:"cd_search_cookies"`
	CD_Cookies        string `json:"cd_cookies"`
	CD_Token_B_T      string `json:"cd_token_bt"`
	Resource_URL      string `json:"resource_url"`
	TripPriceParams   TripPriceParams
	CD_Dispatch_ID    string `json:"cd_dispatch_id"`
}
