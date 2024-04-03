package entities

type Expense struct {
	ID           string  `json:"id"`
	Name 	   string  `json:"name"`
	Description  string  `json:"description"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	DateCreation int64   `json:"date_creation"`
}
