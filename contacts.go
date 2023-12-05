package models

import (
	"database/sql"
	"sync"
)

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

type AddContactParams struct {
	Add_From             string
	Company_Id           string
	Unique_Id            string
	Account_Type         string
	Account_Id           string //For adding pickup/dropoff and driver
	Email                string
	Firstname            string
	Lastname             string
	Fullname             string
	Company_Name         string
	Contact_Info         interface{}
	Insurance_Info       interface{}
	Commission_Structure string
	Commission_Amount    string
	Tax_Id               string
	CD_Carrier_Id        string
	CD_Carrier_Sub_Id    string
	MC_Number            string
	Account_Info         []byte //For error log via Import
	Status               string
	Date                 string
	File_Id              int //via Import
	Data_Type            string
	Comments             string //For logging import error
	Notes                string //For adding pickup/dropoff
	Phone_1              string //For Adding Drivers
	TX                   *sql.Tx
}

type UpdateContactParams struct {
	Company_Id           string
	Account_Type         string
	Account_Id           string
	Contact_Id           string
	Email                string
	Firstname            string
	Lastname             string
	Fullname             string
	Company_Name         string
	Contact_Info         interface{}
	Insurance_Info       interface{}
	Commission_Structure string
	Commission_Amount    string
	Tax_Id               string
	MC_Number            string
	Account_Info         []byte //For error log via Import
	Status               string
	File_Id              int //via Import
	Data_Type            string
	Comments             string //For logging import error
	Notes                string //For adding pickup/dropoff
	Phone_1              string //For Adding Drivers
	TX                   *sql.Tx
}

type ListsContactParams struct {
	Company_Id       int
	Fields           string
	Account_Type     string
	StartFrom        int
	Limit            int
	Dispatch_Through string //For Platform Carriers
}

type ContactOrdersParams struct {
	Company_Id   int
	Account_Type string
	Account_Id   string
	Unique_Id    string
	Fields       []interface{}
	Page_Size    int
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type ContactsOrdersResp struct {
	Orders         []interface{} `json:"order_lists"`
	Total_Returned int           `json:"total_returned"`
}

type ContactDetailsResponse struct {
	ContactDetails interface{} `json:"contact_details"`
	Orders         interface{} `json:"orders"`
	Messages       interface{} `json:"messages"`
	Call_Logs      interface{} `json:"call_logs"`
	Contacts       interface{} `json:"contacts"`
	Files          interface{} `json:"files"`
	Block_History  interface{} `json:"block_history"`
}

type ContactMsgParams struct {
	Company_Id   int
	Account_Type string
	Account_Id   string
	Unique_Id    string
	Fields       []interface{}
	Page_Size    int
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type ContactsMsgsResp struct {
	Messages       []interface{} `json:"message_lists"`
	Total_Returned int           `json:"total_returned"`
}

type ContactCallLogsParams struct {
	Company_Id   int
	Account_Type string
	Account_Id   string
	Unique_Id    string
	Phone_1      string
	Phone_2      string
	Fields       []interface{}
	Page_Size    int
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type ContactsCallLogsResp struct {
	Call_Logs      []interface{} `json:"log_lists"`
	Total_Returned int           `json:"total_returned"`
}

type ContactSubContactsParams struct {
	Company_Id   int
	Account_Type string
	Account_Id   string
	Unique_Id    string
	Fields       []interface{}
	Page_Size    int
	Added_Driver int64 //For Adding driver to cArrier contActs only
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type AccountsContactsResp struct {
	Contacts       []interface{} `json:"contact_lists"`
	Total_Returned int           `json:"total_returned"`
}

type LoadContactFileParams struct {
	Company_Id   int
	Account_Type string
	Account_Id   string
	Page_Size    int
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type ContactsFilesResp struct {
	Files          []interface{} `json:"file_lists"`
	Total_Returned int           `json:"total_returned"`
}

type LoadBlockHistParams struct {
	Company_Id   int
	Account_Type string
	Account_Id   string
	Page_Size    int
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type BlockHistoryResp struct {
	Block_History  []interface{} `json:"history_lists"`
	Total_Returned int           `json:"total_returned"`
}

type ContactStatusParams struct {
	Company_Id   int
	Account_Id   string
	Ref_Order_Id string
	User_Id      int
	Status       string
	Notes        string
	Agent_Name   string
}

type ListPlatformCarrParams struct {
	Company_Id       int
	Dispatch_Through string
}

type FormContactParams struct {
	Company_Id   int
	Account_Id   string
	Account_Type string
	Driver_Id    string
	Carrier_Id   string
	Contact_Id   string
}
