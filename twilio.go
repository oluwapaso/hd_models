package models

import "database/sql"

type AddCallHistoryParams struct {
	Agent_Id             int
	Company_Id           int
	Recording_SID        string
	Recording_Url        string
	Recording_Status     string
	Call_SID             string
	Recording_Start_Time string
	Account_SID          string
	Recording_Duration   string
	From_Number          string
	To_Number            string
	Item_Type            string
	Item_Id              int
	Date                 string
	Tx                   *sql.Tx
}
