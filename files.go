package models

type LIstFilesParams struct {
	Type       string
	Company_Id int
	Item_Type  string
	Item_Id    int
}

type FileLists struct {
	File_Id        int    `json:"file_id"`
	Company_Id     int    `json:"company_id"`
	Item_Type      string `json:"item_type"`
	Item_Id        int    `json:"item_id"`
	Item_Unique_Id string `json:"item_unique_id"`
	File_Name      string `json:"file_name"`
	File_Type      string `json:"file_type"`
	File_Loc       string `json:"file_loc"`
	Extension      string `json:"extension"`
}
