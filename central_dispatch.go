package models

import "database/sql"

type Check_CD_HD_Carrier_Params struct {
	CD_Carrier_Id  string
	CD_Carrier_UId string
}

type Add_CD_Saved_Carrier_Params struct {
	Username     string
	Carrier_Id   string
	Carrier_UId  string
	Company_Name string
	Auth_State   string
	MC_Num       string
	Contact      string
	Address      string
	City         string
	State        string
	Zip          string
	Phone_1      string
	Phone_2      string
	Fax          string
	Tx           *sql.Tx
}

type Add_CD_Cookies_Params struct {
	Company_Id          int
	Cookies             string
	Cookie_Type         string
	Local_Storage_Value string
	B_T                 string
	Tx                  *sql.Tx
}
