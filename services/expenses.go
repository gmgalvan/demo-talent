package services

import (
	"context"
	"fmt"
	"time"

	"github.com/demo-talent/entities"
	"github.com/demo-talent/repository"
)

// ExpenseService defines the interface for expense-related operations.
type ExpenseService interface {
	CreateExpense(ctx context.Context, e *entities.Expense) error
	GetExpenseByID(ctx context.Context, id string) (*entities.Expense, error)
	UpdateExpense(ctx context.Context, e *entities.Expense) error
	DeleteExpense(ctx context.Context, id string) error
}

type expenseServiceImpl struct {
	repo repository.ExpenseRepositoryInterface
}

// NewExpenseService creates a new instance of ExpenseService.
func NewExpenseService(repo repository.ExpenseRepositoryInterface) ExpenseService {
	return &expenseServiceImpl{repo: repo}
}

// CreateExpense creates a new expense.
func (s *expenseServiceImpl) CreateExpense(ctx context.Context, e *entities.Expense) error {
	e.ID = generateUniqueID()

	e.DateCreation = time.Now().Unix()

	return s.repo.Create(ctx, e)
}

// GetExpenseByID retrieves an expense by its ID.
func (s *expenseServiceImpl) GetExpenseByID(ctx context.Context, id string) (*entities.Expense, error) {
	return s.repo.GetByID(ctx, id)
}

// UpdateExpense updates an existing expense.
func (s *expenseServiceImpl) UpdateExpense(ctx context.Context, e *entities.Expense) error {
	_, err := s.repo.GetByID(ctx, e.ID)
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, e)
}

// DeleteExpense deletes an expense by its ID.
func (s *expenseServiceImpl) DeleteExpense(ctx context.Context, id string) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}

// generateUniqueID generates a new unique ID for an expense.
func generateUniqueID() string {
	return fmt.Sprintf("expense_%d", time.Now().UnixNano())
}
