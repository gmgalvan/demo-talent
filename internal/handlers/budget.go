package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/demo-talent/internal/entities"
	"github.com/demo-talent/internal/services"
	"github.com/demo-talent/logger"
	"github.com/gorilla/mux"

)

// BudgetHandlers is the interface for the budgets handlers.
type BudgetHandlers interface {
	CreateBudget() http.HandlerFunc
	GetBudgetByID() http.HandlerFunc
	UpdateBudget() http.HandlerFunc
	DeleteBudget() http.HandlerFunc
	ListBudgets() http.HandlerFunc
} 

// imple is a struct that holds the dependencies for the budget handlers.
type butgetRouter struct { 
	service services.BudgetServiceInterface
	log     *logger.Logger
}
 
// NewBudgetHandlers creates a new instance of BudgetHandlers.
func NewBudgetHandlers(ctx context.Context, service services.BudgetServiceInterface) BudgetHandlers {
	log, ok := ctx.Value("logger").(*logger.Logger)
	if !ok {
		log = logger.NewLogger(false, logger.INFO)
	}
	return &butgetRouter{
		service: service,
		log:     log,
	}
}

// CreateBudget is the HTTP handler for creating a new budget.
func (rb *butgetRouter) CreateBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b entities.Budget
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			message := fmt.Sprintf("Failed to decode request body on create budget: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		if err := rb.service.Create(ctx, &b); err != nil {
			message := fmt.Sprintf("Failed to create budget: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// GetBudgetByID is the HTTP handler for getting a budget by its ID.
func (rb *butgetRouter) GetBudgetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		b, err := rb.service.GetByID(ctx, id)
		if err != nil {
			message := fmt.Sprintf("Failed to get budget by ID: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(b)
	}
}

// UpdateBudget is the HTTP handler for updating a budget.

func (rb *butgetRouter) UpdateBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b entities.Budget
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			message := fmt.Sprintf("Failed to decode request body on update budget: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		if err := rb.service.Update(ctx, &b); err != nil {
			message := fmt.Sprintf("Failed to update budget: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteBudget is the HTTP handler for deleting a budget.
func (rb *butgetRouter) DeleteBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		if err := rb.service.Delete(ctx, id); err != nil {
			message := fmt.Sprintf("Failed to delete budget: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// ListBudgets is the HTTP handler for listing budgets.
func (rb *butgetRouter) ListBudgets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			message := fmt.Sprintf("Failed to get limit query parameter: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			message := fmt.Sprintf("Failed to get offset query parameter: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		budgets, err := rb.service.List(ctx, limit, offset)
		if err != nil {
			message := fmt.Sprintf("Failed to list budgets: %s", err.Error())
			rb.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", message)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(budgets)
	}

}

//swagger:parameters getBudgetRequest deleteBudgetRequest
type budgetIDParameter struct {
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters updateBudgetRequest
type updateBudgetRequest struct {
	// in: body
	Body struct {
		// required: true
		ID string `json:"id"`
		// required: true
		Description string `json:"description"`
		// required: true
		Amount float64 `json:"amount"`
		// required: true
		StartDate time.Time `json:"start_date"`
		// required: true
		EndDate time.Time `json:"end_date"`
	}
}

// swagger:response budgetResponse
type budgetResponse struct {
	// in: body
	Body struct {
		ID		  string    `json:"id"`
		Description string    `json:"description"`
		Amount      float64   `json:"amount"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}
}



