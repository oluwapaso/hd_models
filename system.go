package models

type SystemReports struct {
	Report_Id   int    `json:"report_id"`
	Company_Id  int    `json:"company_id"`
	Agent_Id    int    `json:"agent_id"`
	Report_Type string `json:"report_type"`
	Page        string `json:"page"`
	Message     string `json:"message"`
	Date        string `json:"date"`
}
