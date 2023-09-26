package models

type CitySearchParams struct {
	Account_Id string
	API_Key    string
	Keyword    string
}

type CityExtended struct {
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	County    string `json:"county,omitempty"`
}
