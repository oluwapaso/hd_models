package models

import "sync"

type AddConstExpensesParams struct {
	Company_Id   int
	Name         string
	Amount       string
	Descriptions string
}

type ListExpensesParams struct {
	Company_Id int
	StartFrom  int
	Limit      int
}

type ExpensesList struct {
	Expenses_Id int    `json:"expenses_id"`
	Company_Id  int    `json:"company_id"`
	Duration    string `json:"duration"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
}

type LoadReportsParams struct {
	Company_Id              int
	Data_For                string
	From_Date               string
	To_Date                 string
	Report_Type             string
	Historical_Statuses     []interface{}
	Estimated_Dates_Filter  string
	Define_Orders_As        string
	Include_Orders_Not_Disp string
	Wg                      *sync.WaitGroup
	Ch                      chan interface{}
}

type DataCountsResponse struct {
	PendingTasks             interface{} `json:"pending_tasks"`
	CompletedTasks           interface{} `json:"completed_tasks"`
	TerminatedTasks          interface{} `json:"terminated_tasks"`
	TasksToday               interface{} `json:"tasks_today"`
	ActiveLeads              interface{} `json:"active_leads"`
	ReceivedLeads            interface{} `json:"received_leads"`
	ActiveQuotes             interface{} `json:"active_quotes"`
	ReceivedQuotes           interface{} `json:"received_quotes"`
	FollowUpToday            interface{} `json:"follow_up_today"`
	ActiveOrders             interface{} `json:"active_orders"`
	CreatedOrders            interface{} `json:"created_orders"`
	PostedToLoadboard        interface{} `json:"posted_to_loadboards"`
	PostedToLoadboardInRange interface{} `json:"posted_to_loadboards_in_range"`
	DispatchedOrders         interface{} `json:"dispatched_orders"`
	DispatchedOrdersInRange  interface{} `json:"dispatched_orders_in_range"`
	DeliveredOrders          interface{} `json:"delivered_orders"`
	DeliveredOrdersInRange   interface{} `json:"delivered_orders_in_range"`
	PickedUpOrders           interface{} `json:"pickedup_orders"`
	OrdersWithIssues         interface{} `json:"orders_with_issues"`
	PickupToday              interface{} `json:"pickup_today"`
	DeliveriesToday          interface{} `json:"deliveries_today"`
	LeadsToQuotes            interface{} `json:"leads_to_quotes"`
	LeadsToOrders            interface{} `json:"leads_to_orders"`
	QuotesToOrders           interface{} `json:"quotes_to_orders"`
	SentEmails               interface{} `json:"sent_emails"`
	DeliveredEmails          interface{} `json:"delivered_emails"`
	ErroredEmails            interface{} `json:"errored_emails"`
	UniqueEmailOpens         interface{} `json:"unique_email_opens"`
	UniqueEmailClicks        interface{} `json:"unique_email_clicks"`
	TotalEmailOpens          interface{} `json:"total_email_opens"`
	TotalEmailClicks         interface{} `json:"total_email_clicks"`
	SentSMS                  interface{} `json:"sent_sms"`
	SentSMSViaQueue          interface{} `json:"sent_sms_via_queue"`
}

type PaymentsDataResponse struct {
	ShipperToBroker        interface{} `json:"shipper_to_broker"`
	ShipperToCarrier       interface{} `json:"shipper_to_carrier"`
	BrokerToCarrier        interface{} `json:"broker_to_carrier"`
	CarrierToBroker        interface{} `json:"carrier_to_broker"`
	TotalPaymentsCounts    int         `json:"total_payment_counts"`
	TotalPaymentsReceived  float64     `json:"total_payment_received"`
	TotalPayouts           float64     `json:"total_payouts"`
	TotalPaymentsProcessed float64     `json:"total_payment_processed"`
}

type HistoricalStatusResponse struct {
	Order_Id         int         `json:"order_id"`
	Order_Uniq_Id    string      `json:"order_uniq_id"`
	Dates            interface{} `json:"dates"`
	Assigned_To      int         `json:"assigned_to"`
	Origin           interface{} `json:"origin"`
	Destination      interface{} `json:"destination"`
	Dispatch_Details interface{} `json:"dispatch_details"`
	Vehicles         interface{} `json:"vehicles"`
	Shipper          interface{} `json:"shipper"`
	Payments         interface{} `json:"payments"`
	Type             string      `json:"type"`
	Status           string      `json:"status"`
	Table_Name       string      `json:"table_name"`
}

type EstimatedDatesResponse struct {
	Order_Id         int         `json:"order_id"`
	Order_Uniq_Id    string      `json:"order_uniq_id"`
	Assigned_To      int         `json:"assigned_to"`
	Dates            interface{} `json:"dates"`
	Origin           interface{} `json:"origin"`
	Destination      interface{} `json:"destination"`
	Dispatch_Details interface{} `json:"dispatch_details"`
	Vehicles         interface{} `json:"vehicles"`
	Shipper          interface{} `json:"shipper"`
	Status           string      `json:"status"`
}

type CommmonDealsReportData struct {
	Numb_Of_Reports           int     `json:"num_of_results"`
	TotalLeads                int     `json:"all_total_leads"`
	TotalQuotes               int     `json:"all_total_quotes"`
	TotalOrders               int     `json:"all_total_orders"`
	TotalLeadsAmount          float64 `json:"all_total_leads_amount"`
	TotalDispOrders           int     `json:"all_dispatched_orders"`
	TotalLeadsConvRate        float64 `json:"all_leads_conversion_rate"`
	TotalQuotesConvRate       float64 `json:"all_quotes_conversion_rate"`
	TotalProfitMargin         float64 `json:"all_profit_margin"`
	DispOrderTotalTariff      float64 `json:"all_total_tariff"`
	DispOrderTotalCarrierPay  float64 `json:"all_total_carrier_pay"`
	DispOrderTotalBrokersFee  float64 `json:"all_total_broker_fee"`
	DispOrderTotalGrossProfit float64 `json:"all_gross_profit"`
	TotalAvrgProfitPerLead    float64 `json:"all_avg_profit_per_lead"`
}
type LeadSourceDataRespose struct {
	CommmonDealsReportData
	LeadSourceReportData []LeadSourceReportData `json:"source_list"`
}

type ReferralDataResponse struct {
	CommmonDealsReportData
	ReferralReportData []ReferralsReportData `json:"referral_list"`
}

type PayableReceivableLists struct {
	OrderLists
	Payment_From_To string      `json:"payment_from_to"`
	Amount          interface{} `json:"amount"`
}

type AgentsDataResponse struct {
	CommmonDealsReportData
	Total_Comm_Amount float64            `json:"all_total_comm_amount"`
	AgentsReportData  []AgentsReportData `json:"agents_list"`
}

type PayableRcvbleDataResponse struct {
	TotalOrders           int         `json:"total_orders_count"`
	TotalTariff           interface{} `json:"total_tariff"`
	TotalRcvdFromShipper  interface{} `json:"total_received_from_shipper"`
	TotalBrokersFee       interface{} `json:"total_brokers_fee"`
	BalFromShipper        interface{} `json:"balance_from_shipper"`
	BalToCarrier          interface{} `json:"balance_to_carrier"`
	BalFromCarrier        interface{} `json:"balance_from_carrier"`
	TotalCarrToBrkr       interface{} `json:"carrier_to_broker"`
	TotalOpenBal          interface{} `json:"total_open_balance"`
	TotalPaymentsReceived interface{} `json:"total_payment_rcvd"`
}

type PaymentsAndExpensesDataResponse struct {
	PlanPayments     interface{} `json:"plan_payments"`
	CreditsPayments  interface{} `json:"credits_payments"`
	AgentsPayments   interface{} `json:"agents_payments"`
	MonthlyExpenses  interface{} `json:"monthly_expenses"`
	ConstantExpenses interface{} `json:"constant_expenses"`
}

type LoadReportResponse struct {
	DataCountsResponse       interface{} `json:"data_counts"`
	PaymentsDataResponse     interface{} `json:"payments_data"`
	PaymentsList             interface{} `json:"payments_list"`
	HistoricalStatusResponse interface{} `json:"historical_status"`
	EstimatedDatesResponse   interface{} `json:"estimated_dates"`
	LeadSourceDataRespose    interface{} `json:"lead_source_data"`
	ReferralDataResponse     interface{} `json:"referral_data"`
	AgentsDataResponse       interface{} `json:"agents_data"`
	QuotesData               interface{} `json:"quotes_data"`
	PayableReceivableData    interface{} `json:"payable_recievable_data"`
	PayableReceivableLists   interface{} `json:"payable_recievable_lists"`
	HD_Payments_And_Expenses interface{} `json:"hd_payments_and_expenses"`
}

type AddMonthlyExpensesParams struct {
	Company_Id   int
	Name         string
	Amount       string
	Descriptions string
	Date         string
}

type ListMonthlyExpensesParams struct {
	Company_Id int
	StartFrom  int
	Limit      int
	Year       string
	Month      string
}

type MonthlyExpensesList struct {
	Expenses_Id int    `json:"expenses_id"`
	Company_Id  int    `json:"company_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
}

type TotalConstExpensesParams struct {
	Company_Id int
}
