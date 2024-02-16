package models

import (
	"database/sql"
	"sync"
)

type CommonDealFields struct {
	Company_id           int            `json:"company_id"`
	Car_run              string         `json:"car_run"`
	Ship_via             string         `json:"ship_via"`
	Assigned_to          int            `json:"assigned_to"`
	Original_assigned_to sql.NullInt16  `json:"original_assigned_to"`
	Source               sql.NullString `json:"source"`
	Source_type          sql.NullString `json:"source_type"`
	Shipper              sql.NullString `json:"shipper"`
	Referral             sql.NullString `json:"referral"`
	Origin               sql.NullString `json:"origin"`
	Destination          sql.NullString `json:"destination"`
	Vehicles             sql.NullString `json:"vehicles"`
	Payments             sql.NullString `json:"payments"`
	Type                 sql.NullString `json:"type"`
	Original_email       sql.NullString `json:"original_email"`
}

type OrderFields struct {
	Order_id      int            `json:"order_id"`
	Order_uniq_id string         `json:"order_uniq_id"`
	Mode          sql.NullString `json:"mode"`
}

type QuoteFields struct {
	Quote_id      int            `json:"quote_id"`
	Quote_uniq_id string         `json:"quotet_uniq_id"`
	Quote_token   sql.NullString `json:"quote_token"`
}

type Shipper struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone_1       string `json:"phone_1"`
	Phone_2       string `json:"phone_2,omitempty"`
	Country       string `json:"country,omitempty"`
	Company       string `json:"company,omitempty"`
	Company_Phone string `json:"company_phone,omitempty"`
	Shipper_note  string `json:"shipper_note,omitempty"`
	Address       string `json:"address,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Zip           string `json:"zip,omitempty"`
	Shipper_Id    string `json:"shipper_id"`
}

type CommonDates struct {
	Received             string `json:"received"`
	Quoted               string `json:"quoted"`
	First_available_date string `json:"first_available_date"`
	Imported             string `json:"imported,omitempty"`
	DealsDates
}

type DealsDates struct {
	Next_followup_date     string `json:"next_followup_date,omitempty"`     //Quotes only
	Last_updated           string `json:"last_updated,omitempty"`           //Quotes only
	Last_updated_timestamp string `json:"last_updated_timestamp,omitempty"` //Quotes only
	Last_bomb              string `json:"last_bomb,omitempty"`              //Quotes only
	Created                string `json:"created,omitempty"`                //Orders only
}

type Origin struct {
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	Full_origin string `json:"full_origin"`
}

type Destination struct {
	City             string `json:"city"`
	State            string `json:"state"`
	Zip              string `json:"zip"`
	Location         string `json:"location"`
	Address          string `json:"address"`
	Full_destination string `json:"full_destination"`
}

type DealPayments struct {
	Total_tariff      string `json:"total_tariff"`
	Total_brokers_fee string `json:"total_brokers_fee"`
	Carrier_pay       string `json:"carrier_pay"`
}

type Vehicles struct {
	Vehicle_id      int    `json:"vehicle_id"`
	Year            int    `json:"year"`
	Make            string `json:"make"`
	Model           string `json:"model"`
	Quantity        int    `json:"quantity"`
	Type            string `json:"type"`
	Running         string `json:"running"`
	Ship_Via        string `json:"ship_via"`
	Color           string `json:"color"`
	VIN             string `json:"vin"`
	State           string `json:"state"`
	Lot_Number      string `json:"lot_number"`
	Brokers_Fee     string `json:"brokers_fee"`
	Plate_Number    string `json:"plate_number"`
	Total_Tariff    string `json:"total_tariff"`
	Additional_Info string `json:"additional_info"`
	Full_vehicle    string `json:"full_vehicle"`
}

type Issues struct {
	Id       int
	Date     string
	Type     string
	Notes    string
	Added_By string
}

type ItemTypeCountsParams struct {
	Company_Id string
	Item_Id    int
	Item_Type  string
	Field      string
	Value      int
}

type CountsNotes struct {
	Company_Id int
	Item_Id    int
	Item_Type  string
	Note_Type  string
}

type DealVehicles struct {
	Vehicle_Id      int         `json:"vehicle_id"`
	Vin             string      `json:"vin"`
	Make            string      `json:"make"`
	Year            interface{} `json:"year"`
	Color           string      `json:"color"`
	Model           string      `json:"model"`
	Quantity        int         `json:"quantity"`
	Type            string      `json:"type"`
	State           string      `json:"state"`
	Tariff          string      `json:"tariff"`
	Running         string      `json:"running"`
	Ship_Via        string      `json:"ship_via"`
	Lot_Number      string      `json:"lot_number"`
	Brokers_Fee     string      `json:"brokers_fee"`
	Full_Vehicle    string      `json:"full_vehicle"`
	Plate_Number    string      `json:"plate_number"`
	Total_Tariff    string      `json:"total_tariff"`
	Additional_Info string      `json:"additional_info"`
}

type DeleteDealsParams struct {
	Tx         *sql.Tx
	Item_Type  string
	Item_Id    []interface{}
	Company_Id int
}

type DeleteDealsExtrasParams struct {
	Tx          *sql.Tx
	Extral_Type string
	Item_Type   string
	Item_Id     []interface{}
	Company_Id  int
}

type ChangedDtlsParams struct {
	Item_Type   string
	Posted_Info map[string]interface{}
	Old_Info    DealsDetails
	Target      string
}

type DealsDetails struct {
	Order_Details OrderDetails
	Quote_Details QuoteDetails
	Lead_Details  LeadDetails
}

type LoadDispatchParams struct {
	Company_Id    int
	HD_Post_Id    string
	HD_Disp_Token string
	Fields        string
}

type DispatchLists struct {
	Dispatch_Sheet_Id   int         `json:"dispatch_sheet_id"`
	Company_Id          string      `json:"company_id"`
	Order_Id            string      `json:"order_id"`
	Order_Uniq_Id       string      `json:"order_uniq_id"`
	User_Id             string      `json:"user_id"`
	HD_Disp_Token       string      `json:"hd_disp_token"`
	Disp_Order_Info     interface{} `json:"dispatched_order_info"`
	Carrier_Info        interface{} `json:"carrier_info"`
	Status              string      `json:"status"`
	Date                interface{} `json:"date"`
	Last_Updated        string      `json:"last_updated"`
	Last_Update_By      string      `json:"last_update_by"`
	Cancelled_By        string      `json:"cancelled_by"`
	Cancellation_Reason string      `json:"cancellation_reason"`
	Dispatch_Terms      string      `json:"dispatch_terms"`
	Carrier_Signature   string      `json:"carrier_signature"`
	Status_History      interface{} `json:"status_history"`
	Transit_Status      string      `json:"transit_status"`
}

type HiddenDispAddress struct {
	City  string `json:"city"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}

type RejectDispParams struct {
	Company_Id    int
	HD_Post_Id    string
	HD_Disp_Token string
	Type          string
	Reason        string
	Tx            *sql.Tx
}

type AcceptDispParams struct {
	Company_Id    int
	HD_Post_Id    string
	HD_Disp_Token string
	Signature     string
	Tx            *sql.Tx
}

type UndispatchParams struct {
	Company_Id    int
	HD_Post_Id    string
	HD_Disp_Token string
	Notes         string
	Tx            *sql.Tx
}

type GeneratePDF struct {
	Company_Id    int
	HD_Post_Id    string
	Account_Id    string
	HD_Disp_Token string
	Order_Id      string
}

type GeneratePDFResponse struct {
	Page_Link        string `json:"page_link"`
	PDF_Link         string `json:"pdf_link"`
	Response_Message string `json:"response_message"`
}

type UpDispTransitStatus struct {
	Company_Id     int
	HD_Post_Id     string
	HD_Disp_Token  string
	Status_Date    string
	Transit_Status string
	Tx             *sql.Tx
}

type DispOrderInfo struct {
	Shipper          interface{} `json:"shipper"`
	Origin           interface{} `json:"origin"`
	Destination      interface{} `json:"destination"`
	Dates            interface{} `json:"dates"`
	Vehicles         interface{} `json:"vehicles"`
	Payments         interface{} `json:"payments"`
	Shippers_Notes   string      `json:"shippers_notes"`
	Include_Comments string      `json:"include_comments"`
	Load_On          string      `json:"load_on"`
	Deliver_On       string      `json:"deliver_on"`
}

type AddDispSheetParams struct {
	Company_Id      int
	Order_Id        int
	Order_Uniq_Id   string
	User_Id         int
	HD_Disp_Token   string
	Disp_Order_Info interface{}
	Carrier_Info    interface{}
	Disp_Terms      string
	Tx              *sql.Tx
}

type Undispatch_HD_Order_Params struct {
	Company_Id               int
	Order_Id                 int
	Order_Uniq_Id            string
	User_Id                  int
	HD_Disp_Token            string
	Old_Disp_Info            interface{}
	Agent_Name               string
	Agent_Email              string
	Agent_Forwarding_Address string
	DBCon                    *sql.DB
}

type ReturnDispParams struct {
	DBConn        *sql.DB
	Company_Id    int
	Order_Id      int
	Old_Disp_Info interface{}
}

type Undisp_HD_Order_Params struct {
	Company_Id               int
	Order_Id                 int
	Order_Uniq_Id            string
	User_Id                  int
	HD_Disp_Token            string
	Disp_Info                interface{}
	Agent_Name               string
	Agent_Email              string
	Agent_Forwarding_Address string
	Tx                       *sql.Tx
}

type GetItemCountsParams struct {
	Item_Type  string
	Item_Id    int
	Company_Id int
}

type UpdateCountParams struct {
	Item_Type  string
	Item_Id    int
	Company_Id int
	Field      string
	Value      string
}

type RemoveAutomParams struct {
	Remove_From_Leads  string
	Remove_From_Quotes string
	Remove_From_Orders string
	Company_Id         int
	Automation_Id      int
	Tx                 *sql.Tx
}

type CountDealFilesParams struct {
	Item_Type  string
	Item_Id    int
	Company_Id int
	Tx         *sql.Tx
}

type LoadNextPrevDealsParams struct {
	Company_Id   int
	Agent_Id     int
	Item_Type    string
	Curr_Item_Id int
	Sort_By      string
	Sort_Dir     string
	Wg           *sync.WaitGroup
	Ch           chan interface{}
}

type MarkDealAsSeenParams struct {
	Table      string
	Field      string
	Item_Id    string
	Company_Id int
	Tx         *sql.Tx
}

type UpDealPaymentJsonFieldParams struct {
	Table      string
	Field      string
	Json_Field string
	Item_Id    string
	Company_Id int
	Value      string
	Tx         *sql.Tx
}

type UpDealPaymentsParams struct {
	Table      string
	Field      string
	Item_Id    string
	Company_Id int
	Value      string
	Tx         *sql.Tx
}

type LoadDealFieldParams struct {
	Table      string
	To_Sel     string
	Qry_Field  string
	Qry_Value  string
	Company_Id int
}

type UpDealVehicleJsonFieldParams struct {
	Table       string
	Up_By_Field string
	Extra_Col   string
	Item_Id     string
	Company_Id  int
	Qry_Value   []interface{}
	Tx          *sql.Tx
}

type AddNewVehicleParams struct {
	Table           string
	Up_By_Field     string
	Item_Id         string
	Company_Id      int
	Num_Of_Vehicles int
	Vehicle         string
	Tx              *sql.Tx
}

type AddDealNotiParams struct {
	Company_Id        int
	Agent_Id          int
	Deal_Type         string
	Deal_UId          string
	Broker_Comp_Name  string
	Broker_Comp_Email string
	Agent_Email       string
	Agent_Name        string
	Tx                *sql.Tx
}

type MarkDealAsNoneDupParams struct {
	Company_Id int
	Deal_Type  string
	Deal_Ids   string
	Tx         *sql.Tx
}

type DeleteDealsForeverParams struct {
	Company_Id int
	Deal_Type  string
	Deal_Ids   string
	Tx         *sql.Tx
}

type ContactSMSMsgDeal struct {
	Company_Id    int
	Contact_Phone string
}

type ContactEmailMsgDeal struct {
	Company_Id    int
	Email_Address string
}

type LoadDealByContactPhoneParams struct {
	Item_Type     string
	Search_By     string
	Company_Id    int
	Contact_Phone string
}

type LoadDealByContactEmailParams struct {
	Item_Type     string
	Search_By     string
	Company_Id    int
	Contact_Email string
}
