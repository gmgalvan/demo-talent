package services

import (
	"context"
	"fmt"
	"time"
	"github.com/demo-talent/logger"
	"github.com/demo-talent/internal/entities"
	"github.com/demo-talent/internal/repository"
)

// ExpenseService defines the interface for expense-related operations.
type ExpenseService interface {
	CreateExpense(ctx context.Context, e *entities.Expense) error
	GetExpenseByID(ctx context.Context, id string) (*entities.Expense, error)
	UpdateExpense(ctx context.Context, e *entities.Expense) error
	DeleteExpense(ctx context.Context, id string) error
	ListExpenses(ctx context.Context, page, limit int) ([]entities.Expense, error)
}

type expenseServiceImpl struct {
	repo repository.ExpenseRepositoryInterface
	log  *logger.Logger
}

// NewExpenseService creates a new instance of ExpenseService.
func NewExpenseService(ctx context.Context, repo repository.ExpenseRepositoryInterface) ExpenseService {
	log, ok := ctx.Value("logger").(*logger.Logger)
    if !ok {
		log = logger.NewLogger(false, logger.INFO)
    }
    return &expenseServiceImpl{
		repo: repo, 
		log:  log,
	}
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
		s.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseService", "Error updating expenses")
		return err
	}

	return s.repo.Update(ctx, e)
}

// DeleteExpense deletes an expense by its ID.
func (s *expenseServiceImpl) DeleteExpense(ctx context.Context, id string) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseService", "Error deleting expenses")
		return err
	}

	return s.repo.Delete(ctx, id)
}

// ListExpenses retrieves a list of expenses with pagination.
func (s *expenseServiceImpl) ListExpenses(ctx context.Context, page, limit int) ([]entities.Expense, error) {
    offset := (page - 1) * limit
    return s.repo.List(ctx, limit, offset)
}

// generateUniqueID generates a new unique ID for an expense.
func generateUniqueID() string {
	return fmt.Sprintf("expense_%d", time.Now().UnixNano())
}
