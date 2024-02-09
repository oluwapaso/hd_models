package models

import "database/sql"

type UpdateAQParams struct {
	Company_Id  int
	Source_Id   int
	Auto_Quote  string
	Default_PC  string
	AQ_Formular string
	Tx          *sql.Tx
}

type Auto_Quote_Formular struct {
	Market_Price_Type    string
	Base_Price           string
	Sign                 string
	Additional_Price     string
	Round_Up             string
	Price_Structure      string
	Additional_Car_Price string
	Num_Of_Vehicles      string
}

type TripPriceParams struct {
	AutoQuoteFormular Auto_Quote_Formular
}
