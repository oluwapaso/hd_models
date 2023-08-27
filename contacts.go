package models

type CompanyAccount struct {
	Comp_account_id      string `json:"comp_account_id"`
	Comp_accnt_unique_id string `json:"comp_accnt_unique_id"`
	Company_id           string `json:"company_id"`
	Account_type         string `json:"account_type"`
	Firstname            string `json:"firstname"`
	Lastname             string `json:"lastname"`
	Fullname             string `json:"fullname"`
	Email                string `json:"email"`
	Company_name         string `json:"company_name"`
	Contacts_info        string `json:"contacts_info"`
	Status               string `json:"status"`
	Date_added           string `json:"date_added"`
}

type CotactInfo struct {
	Phone_1   string `json:"phone_1"`
	Phone_2   string `json:"phone_2"`
	Address_1 string `json:"address_1"`
	Address_2 string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type ContactResp struct {
	RepoRes RepoResponse
	Data    CompanyAccount
}

type ContactExtra struct {
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	CompanyName string `json:"company_name"`
	AccountType string `json:"account_type"`
}
