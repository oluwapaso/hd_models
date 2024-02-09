package models

import "database/sql"

type AddNewDriverFilesParams struct {
	Company_Id int
	Carrier_Id string
	File_Type  string
	File_Name  string
	File_URL   string
	File_EXT   string
	Date       string
	Tx         *sql.Tx
}

type DeleteContFileParams struct {
	Company_Id   int
	Account_Id   string
	Account_Type string
	File_Id      int
	Tx           *sql.Tx
}

type AddDealFilesParams struct {
	Company_Id     int
	Item_Type      string
	Item_Id        string
	Item_Unique_Id string
	File_Name      string
	File_URL       string
	File_EXT       string
	Tx             *sql.Tx
}

type DeleteDealFileParams struct {
	Company_Id int
	Item_Id    string
	Item_Type  string
	File_Id    int
	Tx         *sql.Tx
}

type AddImportedFileParams struct {
	Company_Id      int
	File_Loc        string
	Headers         string
	Values          string
	Imported_Data   string
	Default_Agent   string
	Default_Status  string
	Agent_Id        int
	Total_File_Data int
	Date            string
	Tx              *sql.Tx
}
