package services

import (
	"context"
	"fmt"
	"github.com/demo-talent/internal/entities"
	"github.com/demo-talent/internal/repository"
	"github.com/demo-talent/logger"
	"time"
)

// BudgetServiceInterface is an interface for the budget service.
type BudgetServiceInterface interface {
	Create(ctx context.Context, b *entities.Budget) error
	GetByID(ctx context.Context, id string) (*entities.Budget, error)
	Update(ctx context.Context, b *entities.Budget) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]entities.Budget, error)
}

// BudgetService is a service for budgets.
type BudgetService struct {
	repo repository.BudgetRepositoryInterface
	log  *logger.Logger
}

// NewBudgetService creates a new instance of BudgetService.
func NewBudgetService(ctx context.Context, repo repository.BudgetRepositoryInterface) BudgetServiceInterface {
	log, ok := ctx.Value("logger").(*logger.Logger)
	if !ok {
		log = logger.NewLogger(false, logger.INFO)
	}
	return &BudgetService{
		repo: repo,
		log:  log,
	}
}

// Create saves a new budget in the database.
func (s *BudgetService) Create(ctx context.Context, b *entities.Budget) error {
	b.ID = generateUniqueBudgetID()
	if err := s.repo.Create(ctx, b); err != nil {
		s.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetService", "Error creating budget")
		return fmt.Errorf("error creating budget: %w", err)
	}
	return nil
}

// GetByID retrieves a budget from the database by its ID.
func (s *BudgetService) GetByID(ctx context.Context, id string) (*entities.Budget, error) {
	b, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetService", "Error getting budget")
		return nil, fmt.Errorf("error getting budget: %w", err)
	}
	return b, nil
}

// Update updates a budget in the database.
func (s *BudgetService) Update(ctx context.Context, b *entities.Budget) error {
	s.log.Log(logger.INFO, "/aws/demo-talent", "BudgetService", "Updating budget")
	if err := s.repo.Update(ctx, b); err != nil {
		s.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetService", "Error updating budget")
		return fmt.Errorf("error updating budget: %w", err)
	}
	return nil
}

// Delete removes a budget from the database.
func (s *BudgetService) Delete(ctx context.Context, id string) error {
	s.log.Log(logger.INFO, "/aws/demo-talent", "BudgetService", "Deleting budget")
	if err := s.repo.Delete(ctx, id); err != nil {
		s.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetService", "Error deleting budget")
		return fmt.Errorf("error deleting budget: %w", err)
	}
	return nil
}

// List retrieves a list of budgets from the database.
func (s *BudgetService) List(ctx context.Context, limit, offset int) ([]entities.Budget, error) {
	budgets, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		s.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetService", "Error listing budgets")
		return nil, fmt.Errorf("error listing budgets: %w", err)
	}
	return budgets, nil
}

// generateUniqueID generates a new unique ID for an expense.
func generateUniqueBudgetID() string {
	return fmt.Sprintf("budget_%d", time.Now().UnixNano())
}
