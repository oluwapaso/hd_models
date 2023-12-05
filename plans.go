package models

type PlansLists struct {
	Plan_Id    int         `json:"plan_id"`
	Name       string      `json:"plan_name"`
	Price      interface{} `json:"price"`
	More_Price interface{} `json:"more_price"`
	Capacity   int         `json:"capacity"`
}

type AppLists struct {
	App_Id      int         `json:"app_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       interface{} `json:"price"`
}

type ActivatePlanParams struct {
	Company_Id     int
	Amount         float64
	Ref            string
	Transaction    string
	TrxRef         string
	Status         string
	Activated_Apps []interface{}
	More_Capacity  int
}

type AddCapacityParams struct {
	Company_Id    int
	Amount        float64
	Ref           string
	Transaction   string
	TrxRef        string
	Status        string
	More_Capacity int
}

type BuyCreditsParams struct {
	Company_Id    int
	Amount        float64
	Ref           string
	Transaction   string
	TrxRef        string
	Status        string
	Total_Credits int
}
