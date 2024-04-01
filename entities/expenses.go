package entities

type Expense struct {
	ID           string  `json:"id"`
	Description  string  `json:"description"`
	Amount       float64 `json:"amount"`
	DateCreation int64   `json:"date_creation"`
}
