package services


import (
	"context"
	"fmt"
	"github.com/demo-talent/entities"
	"github.com/demo-talent/repository"
	"time"
)

// BudgetService provides operations on budgets.
type BudgetService interface {
	CreateBudget(ctx context.Context, b *entities.Budget) error
	GetBudgetByID(ctx context.Context, id string) (*entities.Budget, error)
	UpdateBudget(ctx context.Context, b *entities.Budget) error
	DeleteBudget(ctx context.Context, id string) error
}

type budgetService struct {
	repo repository.BudgetRepositoryInterface
}

// NewBudgetService creates a new instance of BudgetService.
func NewBudgetService(repo repository.BudgetRepositoryInterface) BudgetService {
	return &budgetService{repo: repo}
}

// CreateBudget creates a new budget.
func (s *budgetService) CreateBudget(ctx context.Context, b *entities.Budget) error {
	b.ID = GenerateUniqueID("budget")
	budget := &entities.Budget{
		ID:        b.ID,
		Amount:    b.Amount,
		StartDate: b.StartDate,
		EndDate:   b.EndDate,
	}
	if err := s.repo.Create(ctx, budget); err != nil {
		return fmt.Errorf("failed to create budget: %w", err)
	}
	return nil
}

// GetBudgetByID retrieves a budget by its ID.
func (s *budgetService) GetBudgetByID(ctx context.Context, id string) (*entities.Budget, error) {
	budget, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve budget: %w", err)
	}
	return budget, nil
}

// UpdateBudget updates an existing budget.
func (s *budgetService) UpdateBudget(ctx context.Context, b *entities.Budget) error {
	budget := &entities.Budget{
		ID:        b.ID,
		Amount:    b.Amount,
		StartDate: b.StartDate,
		EndDate:   b.EndDate,
	}
	if err := s.repo.Update(ctx, budget); err != nil {
		return fmt.Errorf("failed to update budget: %w", err)
	}
	return nil
}

// DeleteBudget deletes a budget by its ID.
func (s *budgetService) DeleteBudget(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete budget: %w", err)
	}
	return nil
}

func generateUniqueID() string {
	return fmt.Sprintf("budget_%d", time.Now().UnixNano())
}