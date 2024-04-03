package routes

import (
	"github.com/demo-talent/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupExpensesRoutes(r *mux.Router, expensesHandlers handlers.ExpenseHandlers) {
	r.HandleFunc("/expenses", expensesHandlers.CreateExpense()).Methods("POST")
	r.HandleFunc("/expenses/{id}", expensesHandlers.GetExpense()).Methods("GET")
	r.HandleFunc("/expenses", expensesHandlers.UpdateExpense()).Methods("PUT")
	r.HandleFunc("/expenses", expensesHandlers.DeleteExpense()).Methods("DELETE")
	r.HandleFunc("/expenses", expensesHandlers.ListExpenses()).Methods("GET")
}