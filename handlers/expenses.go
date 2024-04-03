package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"context"
	"time"
	"github.com/demo-talent/entities"
	"github.com/demo-talent/services"
	"github.com/demo-talent/logger"
	"github.com/gorilla/mux"
)

type ExpenseHandlers interface {
	CreateExpense() http.HandlerFunc
    GetExpense() http.HandlerFunc
    UpdateExpense() http.HandlerFunc
    DeleteExpense() http.HandlerFunc
    ListExpenses() http.HandlerFunc
}

// ExpenseRouter is the router for the expense handlers.
type expenseRouter struct {
	svc services.ExpenseService
	log  *logger.Logger 
}

func NewExpensesRouter(ctx context.Context, svc services.ExpenseService) ExpenseHandlers {
	log, ok := ctx.Value("logger").(*logger.Logger)
    if !ok {
		log = logger.NewLogger(false, logger.INFO)
    }
	return &expenseRouter{
		svc: svc,
		log: log,
	}
}

// CreateExpense is the HTTP handler for creating a new expense.
// swagger:route POST /expenses Expense createExpenseRequest
// Creates a new expense.
// Responses:
//
//	201: expenseResponse
//	400: errorResponse
//	500: errorResponse
func (rexp *expenseRouter) CreateExpense() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var e entities.Expense
        if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
        defer cancel() 

        if err := rexp.svc.CreateExpense(ctx, &e); err != nil {
            message := fmt.Sprintf("Failed to create expense: %s", err.Error())
			http.Error(w, message, http.StatusInternalServerError)
			//svc.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", message)
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
func (rexp *expenseRouter)GetExpense() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			http.Error(w, "Missing expense ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		expense, err := rexp.svc.GetExpenseByID(ctx, id)
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
func (rexp *expenseRouter)UpdateExpense() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e entities.Expense
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := rexp.svc.UpdateExpense(ctx, &e); err != nil {
			message := fmt.Sprintf("Failed to update expense: %s", err.Error())
			http.Error(w, message, http.StatusInternalServerError)
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
func (rexp *expenseRouter)DeleteExpense() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing expense ID", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		if err := rexp.svc.DeleteExpense(ctx, id); err != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// ListExpenses is the HTTP handler for retrieving a list of expenses with pagination.
// swagger:route GET /expenses Expense listExpensesRequest
// Retrieves a list of expenses with pagination.
// Responses:
//
//	200: expenseResponse
//	400: errorResponse
//	500: errorResponse
func (rexp *expenseRouter)ListExpenses() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        pageStr := r.URL.Query().Get("page")
        limitStr := r.URL.Query().Get("limit")

        if pageStr == "" || limitStr == "" {
            http.Error(w, "Missing page or limit", http.StatusBadRequest)
            return
        }

        page, err := strconv.Atoi(pageStr)
        if err != nil {
            http.Error(w, "Invalid page value", http.StatusBadRequest)
            return
        }

        limit, err := strconv.Atoi(limitStr)
        if err != nil {
            http.Error(w, "Invalid limit value", http.StatusBadRequest)
            return
        }

        ctx := r.Context()
        expenses, err := rexp.svc.ListExpenses(ctx, page, limit)
        if err != nil {
            http.Error(w, "Failed to retrieve expenses", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(expenses)
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
