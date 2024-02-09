package models

import "database/sql"

type AddNotificationsParams struct {
	Company_Id     string `json:"company_id"`
	Item_Id        int    `json:"item_id"`
	Item_Unique_Id string `json:"item_unique_id"`
	Item_Type      string `json:"item_type"`
	Agent_Id       int    `json:"agent_id"`
	Message        string `json:"message"`
	Date_Added     string `json:"date_updated"`
}

type ListsNotificationsParams struct {
	Company_Id    int
	Status        string
	Fields        string
	StartFrom     int
	Limit         int
	View_User     string
	View_Agent_Id string
	Tx            *sql.Tx
}

type MarkNotiAsSeenParams struct {
	Mark_By         string
	Company_Id      int
	Agent_Id        int
	Notification_Id int
	Tx              *sql.Tx
}
