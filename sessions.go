package models

import "database/sql"

type CheckRegTokenParams struct {
	Email string
}

type CheckResetTokenParams struct {
	Email string
}

type ValidateRegTokenParams struct {
	Email string
	Token string
}

type AddTokenParams struct {
	Token_Type string
	Token      string
	Email      string
	Date       string
	Tx         *sql.Tx
}

type DeleteTokenParams struct {
	Token_Type string
	Token      string
	Email      string
	Tx         *sql.Tx
}

type CheckProspectParams struct {
	Load_By string
	Email   string
	Fields  string
}

type UpdateProspectTokenParams struct {
	Token string
	Email string
	Date  string
	Tx    *sql.Tx
}

type AddProspectParams struct {
	Token string
	Email string
	Date  string
	Tx    *sql.Tx
}

type DeleteProsParams struct {
	Token string
	Email string
	Tx    *sql.Tx
}

type CreateCompanyParams struct {
	DepartmentsInfo string
	Email           string
	Date            string
	Tx              *sql.Tx
}

type AddDefaultSettParams struct {
	Company_Id        int
	System_Sent_Email string
	Deal_Tags         string
	Default_Mailer    string
	Tx                *sql.Tx
}

type AddAcntToMappedStParams struct {
	Company_Id int
	Tx         *sql.Tx
}

type AddDefaultTempsParams struct {
	Company_Id int
	Tx         *sql.Tx
}

type DeleteAgentSessionsParams struct {
	Delete_By  string
	Agent_Id   int
	Session_Id int
	Tx         *sql.Tx
}

type AddLoginHistParams struct {
	Company_Id int
	Agent_Id   int
	IP_Address string
	Date       string
	Tx         *sql.Tx
}

type LoginCountParams struct {
	Company_Id int
	Agent_Id   int
	Today      string
}

type LoginCountResp struct {
	Success          bool
	AllNotifications int
	AllTickets       int
	TaskToday        int
	Pending          int
	Finished         int
	Cancelled        int
	AllErrors        int
	NewMessages      int
}

type AddNewSessionParams struct {
	Company_Id          int
	Agent_Id            int
	Device              string
	Access_Token        string
	Access_Token_expiry string
	Last_Login          string
	Tx                  *sql.Tx
}

type AddNewAdminSessionParams struct {
	Admin_Id            int
	Access_Token        string
	Access_Token_expiry string
	Last_Login          string
	Tx                  *sql.Tx
}
