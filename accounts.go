package models

type Accounts struct {
	Account_id        int    `json:"account_id"`
	Company_unique_id string `json:"company_unique_id"`
	Created_by        string `json:"created_by"`
	Company_name      string `json:"company_name"`
	DefaultSettngs
}

type DefaultSettngs struct {
	Settings_id      int `json:"settings_id"`
	Company_id       int `json:"company_id"`
	Sendgrid_api_key int `json:"sendgrid_api_key"`
}

type AccountRepoResp struct {
	Resp RepoResponse
	Data []interface{}
}

type BrokerErrorInfo struct {
	Company_Id        int    `json:"settings_id"`
	Agent_id          int    `json:"agent_id"`
	Agent_name        string `json:"agent_name"`
	Category          string `json:"category"`
	Type              string `json:"type"`
	Complex_message   string `json:"complex_message"`
	Possible_error    string `json:"possible_error"`
	Possible_solution string `json:"possible_solution"`
	Affected_username string `json:"affected_username"`
	Simple_message    string `json:"simple_message"`
	Deal_type         string `json:"deal_type"` //Default None
	Deal_id           string `json:"deal_id"`
	Message_displayed string `json:"message_displayed"`
	Log_for_system    string `json:"log_for_system"` //Default No
}

type ChargesDetails struct {
	Amount         int    `json:"amount"`
	Company_id     int    `json:"company_id"`
	Unique_ids     string `json:"unique_ids"`
	Table_name     string `json:"table_name"`
	Num_Of_Charges int    `json:"num_of_charges"`
}