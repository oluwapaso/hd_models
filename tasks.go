package models

import "database/sql"

type UpdateTaskParams struct {
	Tx           *sql.Tx
	UpdateField  string
	UpdateValues []interface{}
	Where        string
}

type UpdateMultiTaskParams struct {
	Tx       *sql.Tx
	Status   string
	Task_Ids []interface{}
	Agent_Id int
}

type ListsTasksParams struct {
	Company_Id    int
	Fields        string
	StartFrom     int
	Limit         int
	Item_Type     string
	Item_Id       string
	From_Date     string
	To_Date       string
	Priority      string
	View_User     string
	View_Agent_Id string
}

type ListsDealTasksParams struct {
	Company_Id int
	Fields     string
	Item_Type  string
	Item_Id    string
}

type CountTaskParams struct {
	Date       string
	Agent_Id   int
	Company_Id int
}

type AddTaskParams struct {
	Company_Id int
	Item_Type  string
	Item_Id    string
	Agent_Id   int
	Todo       string
	Added_By   string
	Priority   string
	Date       string
}
