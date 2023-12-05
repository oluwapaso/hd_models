package models

type AddNotificationsParams struct {
	Company_Id     string `json:"company_id"`
	Item_Id        int    `json:"item_id"`
	Item_Unique_Id string `json:"item_unique_id"`
	Item_Type      string `json:"item_type"`
	Agent_Id       int    `json:"agent_id"`
	Message        string `json:"message"`
	Date_Added     string `json:"date_updated"`
}
