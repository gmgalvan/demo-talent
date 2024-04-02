package repository


import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/demo-talent/entities"
)

type BudgetRepositoryInterface interface {
	Create(ctx context.Context, b *entities.Budget) error
	GetByID(ctx context.Context, id string) (*entities.Budget, error)
	Update(ctx context.Context, b *entities.Budget) error
	Delete(ctx context.Context, id string) error
}

type BudgetRepository struct {
	db *sql.DB
}

// NewBudgetRepository creates a new instance of BudgetRepository.
func NewBudgetRepository(db *sql.DB) BudgetRepositoryInterface {
	return &BudgetRepository{db: db}
}

// Create saves a new budget in the database.
func (r *BudgetRepository) Create(ctx context.Context, b *entities.Budget) error {
	query := `
		INSERT INTO budgets (id, amount, start_date, end_date)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, query, b.ID, b.Amount, b.StartDate, b.EndDate)
	if err != nil {
		log.Printf("Error creating budget: %v", err)
		return fmt.Errorf("error creating budget: %w", err)
	}
	return nil
}

// GetByID retrieves a budget from the database by its ID.
func (r *BudgetRepository) GetByID(ctx context.Context, id string) (*entities.Budget, error) {
	query := `
		SELECT id, amount, start_date, end_date
		FROM budgets
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var b entities.Budget
	err := row.Scan(&b.ID, &b.Amount, &b.StartDate, &b.EndDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("budget not found with ID: %s", id)
		}
		log.Printf("Error retrieving budget: %v", err)
		return nil, fmt.Errorf("error retrieving budget: %w", err)
	}
	return &b, nil
}

// Update modifies an existing budget in the database.
func (r *BudgetRepository) Update(ctx context.Context, b *entities.Budget) error {
	query := `
		UPDATE budgets
		SET amount = $1, start_date = $2, end_date = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query, b.Amount, b.StartDate, b.EndDate, b.ID)
	if err != nil {
		log.Printf("Error updating budget: %v", err)
		return fmt.Errorf("error updating budget: %w", err)
	}
	return nil
}

// Delete removes a budget from the database by its ID.
func (r *BudgetRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM budgets
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting budget: %v", err)
		return fmt.Errorf("error deleting budget: %w", err)
	}
	return nil
}
