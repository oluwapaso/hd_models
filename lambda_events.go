package models

type MailerEvent struct {
	Server        string `json:"server"`
	From_Email    string `json:"from_email"`
	From_Name     string `json:"from_name"`
	To_Email      string `json:"to_email"`
	To_Name       string `json:"to_name"`
	Subject       string `json:"subject"`
	Email_Body    string `json:"email_body"`
	SMTP_Host     string `json:"smtp_host"`
	SMTP_Port     int    `json:"smtp_port"`
	SMTP_Username string `json:"smtp_username"`
	SMTP_Password string `json:"smtp_password"`
	SendGrid_Key  string `json:"sendgrid_key"`
}
