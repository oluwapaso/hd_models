package models

import "database/sql"

type SysUpdatesParams struct {
	Search_Type string
	Company_Id  int
	Fields      string
	StartFrom   int
	Limit       int
	Year        string
	Type        string
}

type FeatureRequestsParams struct {
	List_Type  string
	Fields     string
	StartFrom  int
	Limit      int
	Agent_Id   int
	Title      string
	Type       string
	Status     string
	Keyword    string
	Sort       string
	Request_Id int
}

type AddFeatureRequestParams struct {
	Company_Id   int
	Agent_Id     int
	Agent_Name   string
	Agent_Email  string
	Company_Name string
	Title        string
	Message      string
	Type         string
	Date         string
	Tx           *sql.Tx
}

type CountReqVotesParams struct {
	Count_Type string
	Agent_Id   int
	Company_Id int
	Request_Id int
	Tx         *sql.Tx
}

type DeleReqVotesParams struct {
	Agent_Id   int
	Company_Id int
	Request_Id int
	Tx         *sql.Tx
}

type UpdateNumOfVotesParams struct {
	Request_Id   int
	Num_Of_Votes int64
	Tx           *sql.Tx
}

type AddReqVotesParams struct {
	Agent_Id   int
	Company_Id int
	Request_Id int
	Vote_Type  string
	Date       string
	Tx         *sql.Tx
}

type TicketListsParams struct {
	Company_Id int
	Agent_Id   int
	Fields     string
	Filter_By  string
	Keyword    string
	StartFrom  int
	Limit      int
}

type UpdateTicketSeenParams struct {
	Seen_By   string
	Ticket_Id int
	Tx        *sql.Tx
}

type TicketDetailsParams struct {
	Company_Id    int
	Agent_Id      int
	Fields        string
	Ticket_Id     int
	Track_Type    string
	Ticket_Number string
}

type TicketRespParams struct {
	Company_Id int
	Agent_Id   int
	Fields     string
	Track_Type string
	Ticket_Id  int
}

type AddTicketResponseParams struct {
	Company_Id    int
	Agent_Id      int
	Admin_Id      int
	Client_Name   string
	Message       string
	Ticket_Id     int
	Response_From string
	RepliedBy     string
	Date          string
	Tx            *sql.Tx
}

type UpdatePingParams struct {
	Ticket_Id int
	RepliedBy string
	Date      string
	Tx        *sql.Tx
}

type AddNewTicketParams struct {
	Company_Id    int
	Agent_Id      int
	Fullname      string
	Phone         string
	Email         string
	Company_Name  string
	Inquiry       string
	Message       string
	Date          string
	Ticket_Number string
	Tx            *sql.Tx
}

type AdminTicketListsParams struct {
	Fields    string
	Filter_By string
	StartFrom int
	Limit     int
}

type UpdateTicketStatusParams struct {
	Ticket_Id int
	Status    string
	Tx        *sql.Tx
}

type UpdateFeatReqStatusParams struct {
	Request_Id int
	Status     string
	Tx         *sql.Tx
}

type FeaturesLookupParams struct {
	Fields  string
	Keyword string
	Limit   int
}

type MergeFeatReqParams struct {
	Request_Id int
	Merge_To   int
	Tx         *sql.Tx
}
