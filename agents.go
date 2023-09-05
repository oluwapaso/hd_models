package models

import "database/sql"

type Agents struct {
	Agent_id         int            `json:"agent_id"`
	Company_id       string         `json:"company_id"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone,omitempty"`
	Lead_multiples   int            `json:"lead_multiples,omitempty"`
	Username         string         `json:"username,omitempty"`
	Mapped_import    string         `json:"mapped_import,omitempty"`
	Webmail_settings sql.NullString `json:"webmail_settings,omitempty"`
	Mailer_settings  sql.NullString `json:"mailer_settings,omitempty"`
}

type MultiAgentsResp struct {
	RepoRes RepoResponse
	Data    []Agents
}
