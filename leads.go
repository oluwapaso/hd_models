package models

type LeadLists struct {
	Lead_Id              int         `json:"lead_id"`
	Lead_Uniq_Id         string      `json:"lead_uniq_id"`
	Company_Id           int         `json:"company_id"`
	Dates                interface{} `json:"dates"`
	Car_Run              string      `json:"car_run"`
	Ship_Via             string      `json:"ship_via"`
	Assigned_To          int         `json:"assigned_to"`
	Original_Assigned_To int         `json:"original_assigned_to"`
	Assigned_To_Info     interface{} `json:"assigned_to_info"`
	Origin               interface{} `json:"origin"`
	Destination          interface{} `json:"destination"`
	Vehicles             interface{} `json:"vehicles"`
	Shipper              interface{} `json:"shipper"`
	Type                 string      `json:"type"`
	Source               string      `json:"source"`
	Source_Type          string      `json:"source_type"`
	Status               string      `json:"status"`
	Item_Counts          interface{} `json:"item_counts"`
	Subscribed_Drips     interface{} `json:"subscribed_drips"`
}

type LeadDetails struct {
	LeadLists
	//Not finished yet 11/22/2023
}
