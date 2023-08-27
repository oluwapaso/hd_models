package models

type Agents struct {
	Agent_id       int    `json:"agent_id"`
	Company_id     string `json:"company_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Lead_multiples int    `json:"lead_multiples"`
	Username       string `json:"username"`
	Mapped_import  string `json:"mapped_import,omitempty"`
}

type AgentResp struct {
	RepoRes RepoResponse
	Data    []Agents
}
