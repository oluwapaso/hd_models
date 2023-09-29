package models

type AddExpensesParams struct {
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
