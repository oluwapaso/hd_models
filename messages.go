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

type IMAP_Event struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
