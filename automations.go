package models

import "database/sql"

type AutomationListsParams struct {
	Search_By  string
	Fields     string
	Company_Id int
	Agent_Id   int
	StartFrom  int
	Limit      int
	Name       string
}

type LoadAutomationsParams struct {
	Search_By  string
	Fields     string
	Company_Id int
	Agent_Id   int
	Used_For   string
	StartFrom  int
	Limit      int
	Name       string
}

type LoadSingleAutomationsParams struct {
	Search_By     string
	Fields        string
	Company_Id    int
	Name          string
	Automation_Id int
	Parent_Id     string
	Order_By      string
	Limit         int
}

type AddAutomationParams struct {
	Company_Id      int
	Agent_Id        int
	Used_For        string
	Name            string
	Date            string
	Default_Trigger interface{}
	Version         int
	Tx              *sql.Tx
}

type AddMultipleStepsParams struct {
	Company_Id    int
	Agent_Id      int
	Automation_Id int
	Step_UID      string
	Step_Type     string
	Parent_Id     string
	Parent_UID    string
	Parent_Type   string
	Step_Position string
	Children      string
	Event_Info    string
}

type UpdateAutomationParams struct {
	Update_Type     string
	Company_Id      int
	Agent_Id        int
	Automation_Id   int
	Automation_Name string
	Parent_Id       int
	Trigger         interface{}
	Assigned_Agents interface{}
	Must_Update     bool
	Status          string
	Is_Published    string
	Tx              *sql.Tx
}

type LoadAutomationsStepsParams struct {
	Company_Id    int
	Automation_Id int
	Fields        string
}

type LoadAutomationVersionParams struct {
	Fields     string
	Company_Id int
	Parent_Id  int
}

type LoadStepDetailsParams struct {
	Load_Type     string
	Fields        string
	Company_Id    int
	Automation_Id int
	Step_Id       int
	Step_UID      string
	Step_Position int
	Parent_Type   string
}

type LoadStepChildrenParams struct {
	Fields        string
	Company_Id    int
	Step_Id       int
	Automation_Id int
}

type AddNewStepParams struct {
	Company_Id    int
	Agent_Id      int
	Automation_Id int
	Tx            *sql.Tx
}

type GetStepPositionParams struct {
	Company_Id    int
	Automation_Id int
	Old_Position  string
	New_Step_Id   int64
	Fields        string
}

type UpdateStepParams struct {
	Update_Type   string
	Company_Id    int
	Agent_Id      int
	Automation_Id int
	Step_Position int
	Step_Id       int
	Must_Update   bool
	New_Step_Id   int64
	New_Step_UId  string
	Old_Step_Id   string
	Step_UID      string
	Parent_Id     string
	Parent_UID    string
	Parent_Type   string
	Children      interface{}
	Event_Info    string
	Tx            *sql.Tx
}

type DeleteStepParams struct {
	Company_Id    int
	Step_Id       int
	Automation_Id int
	Tx            *sql.Tx
}

type DeleteAutomationParams struct {
	Company_Id    int
	Automation_Id int
	Tx            *sql.Tx
}

type DeleteTodoDripsParams struct {
	Delete_Type   string
	Company_Id    int
	Automation_Id int
	Item_Type     string
	Item_Id       int
	Tx            *sql.Tx
}

type DeleteMultipleTodoDripsParams struct {
	Delete_Type string
	Company_Id  int
	Item_Type   string
	Item_Ids    []interface{}
	Tx          *sql.Tx
}

type TriggerAutomParams struct {
	Company_Id      int
	Automation_Type string
	Item_Info       map[string]interface{}
	Tx              *sql.Tx
}

type InitiateCampaignParams struct {
	Company_Id      int
	Automation_Id   int
	Automation_Name string
	Item_Id         int
	Item_Type       string
	Agent_Id        int
	Tx              *sql.Tx
}

type AddAutomLogParams struct {
	Company_Id        int
	Automation_Id     int
	Agent_Id          int
	Item_Id           int
	Item_Type         string
	Event_Name        string
	Message           string
	Step_UID          string
	Date              string
	Log_Type          string
	Exception_Message string
	Tx                *sql.Tx
}

type TodoAutomListsParams struct {
	Fields string
	Limit  int
}

type ProcessNextDripParams struct {
	Company_Id    int
	Automation_Id int
	Item_Id       int
	Item_Type     string
	Step_Position int
	Todo_Id       int
	Agent_Id      int
	Tx            *sql.Tx
}

type DelTodoStepParams struct {
	Company_Id    int
	Agent_Id      int
	Automation_Id int
	Item_Id       int
	Item_Type     string
	Todo_Id       int
	Tx            *sql.Tx
}

type TerminateAutomParams struct {
	Company_Id    int
	Agent_Id      int
	Automation_Id int
	Item_Id       int
	Item_Type     string
	Event_Name    string
	Reason        string
	Todo_Id       string
	Step_Position string
	Step_UID      string
	Delete_Todo   string
	Tx            *sql.Tx
}

type AutoTemplateParams struct {
	Template_Id   int
	Company_Id    int
	Fields        string
	Template_Type string
}
