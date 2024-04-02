package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/demo-talent/entities"
	"github.com/demo-talent/services"
	"time"
)


//CreateBudget is the HTTP handler for creating a new budget.
// swagger:route POST /budgets Budget createBudgetRequest
// Creates a new budget.
// Responses:
//
//	201: budgetResponse
//	400: errorResponse
//	500: errorResponse
func CreateBudget(svc services.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b entities.Budget
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := svc.CreateBudget(ctx, &b); err != nil {
			http.Error(w, "Failed to create budget", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)
	}
}

//GetBudget is the HTTP handler for retrieving a budget by ID.
// swagger:route GET /budgets/{id} Budget getBudgetRequest
// Retrieves a budget by ID.
// Responses:
//
//	200: budgetResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
func GetBudget(svc services.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing budget ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		budget, err := svc.GetBudgetByID(ctx, id)
		if err != nil {
			http.Error(w, "Failed to retrieve budget", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(budget)
	}
}

// UpdateBudget is the HTTP handler for updating an existing budget.
// swagger:route PUT /budgets Budget updateBudgetRequest
// Updates an existing budget.
// Responses:
//
//	200: budgetResponse
//	400: errorResponse
//	500: errorResponse
func UpdateBudget(svc services.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b entities.Budget
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := svc.UpdateBudget(ctx, &b); err != nil {
			http.Error(w, "Failed to update budget", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(b)
	}
}

// DeleteBudget is the HTTP handler for deleting a budget by ID.
// swagger:route DELETE /budgets/{id} Budget deleteBudgetRequest
// Deletes a budget by ID.
// Responses:
//
//	200: okResponse
//	400: errorResponse
//	500: errorResponse
func DeleteBudget(svc services.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing budget ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := svc.DeleteBudget(ctx, id); err != nil {
			http.Error(w, "Failed to delete budget", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Budget represents a budget entity
type Budget struct {
    ID        int       `json:"id"`
    Amount    float64   `json:"amount"`
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}

// swagger:response budgetResponse
type budgetResponse struct {
    // in:body
    Body struct {
        ID        int       `json:"id"`
        Amount    float64   `json:"amount"`
        StartDate time.Time `json:"start_date"`
        EndDate   time.Time `json:"end_date"`
    }
}