package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/demo-talent/entities"
	"github.com/demo-talent/services"
)

// HelloWorld is the HTTP handler for the root path.
// swagger:route GET / HelloWorld helloWorldRequest
// Returns a simple hello world message.
// Responses:
//
//	200: okResponse
//	500: errorResponse
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// CreateExpense is the HTTP handler for creating a new expense.
// swagger:route POST /expenses Expense createExpenseRequest
// Creates a new expense.
// Responses:
//
//	201: expenseResponse
//	400: errorResponse
//	500: errorResponse
func CreateExpense(svc services.ExpenseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e entities.Expense
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := svc.CreateExpense(ctx, &e); err != nil {
			http.Error(w, "Failed to create expense", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)
	}
}

// GetExpense is the HTTP handler for retrieving an expense by ID.
// swagger:route GET /expenses/{id} Expense getExpenseRequest
// Retrieves an expense by ID.
// Responses:
//
//	200: expenseResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
func GetExpense(svc services.ExpenseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing expense ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		expense, err := svc.GetExpenseByID(ctx, id)
		if err != nil {
			http.Error(w, "Failed to retrieve expense", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(expense)
	}
}

// UpdateExpense is the HTTP handler for updating an expense.
// swagger:route PUT /expenses Expense updateExpenseRequest
// Updates an expense.
// Responses:
//
//	200: okResponse
//	400: errorResponse
//	500: errorResponse
func UpdateExpense(svc services.ExpenseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e entities.Expense
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := svc.UpdateExpense(ctx, &e); err != nil {
			http.Error(w, "Failed to update expense", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteExpense is the HTTP handler for deleting an expense by ID.
// swagger:route DELETE /expenses/{id} Expense deleteExpenseRequest
// Deletes an expense by ID.
// Responses:
//
//	200: okResponse
//	400: errorResponse
//	500: errorResponse
func DeleteExpense(svc services.ExpenseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing expense ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := svc.DeleteExpense(ctx, id); err != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// swagger:parameters createExpenseRequest
type createExpenseRequest struct {
	// in:body
	Body struct {
		// Required: true
		Description string `json:"description"`
		// Required: true
		Amount float64 `json:"amount"`
	}
}

// swagger:parameters getExpenseRequest deleteExpenseRequest
type expenseIDParameter struct {
	// in:path
	// Required: true
	ID string `json:"id"`
}

// swagger:parameters updateExpenseRequest
type updateExpenseRequest struct {
	// in:body
	Body struct {
		// Required: true
		ID string `json:"id"`
		// Required: true
		Description string `json:"description"`
		// Required: true
		Amount float64 `json:"amount"`
	}
}

// swagger:response expenseResponse
type expenseResponse struct {
	// in:body
	Body struct {
		ID           string  `json:"id"`
		Description  string  `json:"description"`
		Amount       float64 `json:"amount"`
		DateCreation int64   `json:"date_creation"`
	}
}

// swagger:response errorResponse
type errorResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response okResponse
type okResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}
