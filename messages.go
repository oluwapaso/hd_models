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
	Thread_ID            int            `json:"thread_id"`
	Message_id           string         `json:"message_id"`
	Unique_id            string         `json:"unique_id"`
	Message_type         string         `json:"message_type"`
	Agent_id             string         `json:"agent_id"`
	Deal_type            string         `json:"deal_type"`
	Deal_id              sql.NullString `json:"deal_id"`
	Company_id           string         `json:"company_id"`
	Contact_name         sql.NullString `json:"contact_name"`
	Contact_type         string         `json:"contact_type"`
	Forwarding_email     sql.NullString `json:"forwarding_email"`
	Send_from            string         `json:"send_from"`
	Send_to              string         `json:"send_to"`
	Subject              sql.NullString `json:"subject"`
	Message_body         sql.NullString `json:"message_body"`
	Date                 string         `json:"date"`
	Last_send_attempt    sql.NullString `json:"last_send_attempt"`
	Sent_by              string         `json:"sent_by"`
	Status               string         `json:"status"`
	Outgoing_mail_status string         `json:"outgoing_mail_status"`
	Outgoing_mail_info   sql.NullString `json:"outgoing_mail_info"`
}
