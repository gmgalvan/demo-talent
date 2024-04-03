package routes


import (
	"github.com/demo-talent/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupBudgetRoutes(r *mux.Router, budgetHandlers handlers.BudgetHandlers) {
	r.HandleFunc("/budgets", budgetHandlers.CreateBudget()).Methods("POST")
	r.HandleFunc("/budgets/{id}", budgetHandlers.GetBudgetByID()).Methods("GET")
	r.HandleFunc("/budgets", budgetHandlers.UpdateBudget()).Methods("PUT")
	r.HandleFunc("/budgets", budgetHandlers.DeleteBudget()).Methods("DELETE")
	r.HandleFunc("/budgets", budgetHandlers.ListBudgets()).Methods("GET")
}
