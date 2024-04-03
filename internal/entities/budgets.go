package entities


type Budget struct {
    ID          string       `json:"id"`
    Description string    `json:"description"`
    Amount      float64   `json:"amount"`
    StartDate   string `json:"start_date"`
    EndDate     string `json:"end_date"`
}
