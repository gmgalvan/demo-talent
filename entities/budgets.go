package entities

import "time"

type Budget struct {
    ID        string       `json:"id"`
	// needs desciption and name
    Amount    float64   `json:"amount"`
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}