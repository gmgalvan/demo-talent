package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/demo-talent/internal/entities"
	"github.com/demo-talent/logger"
)

// BudgetRepositoryInterface is an interface for the budget repository.
type BudgetRepositoryInterface interface {
	Create(ctx context.Context, b *entities.Budget) error
	GetByID(ctx context.Context, id string) (*entities.Budget, error)
	Update(ctx context.Context, b *entities.Budget) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]entities.Budget, error)
}

// BudgetRepository is a repository for budgets.
type BudgetRepository struct {
	db  *sql.DB
	log *logger.Logger
}

// NewBudgetRepository creates a new instance of BudgetRepository.
func NewBudgetRepository(ctx context.Context, db *sql.DB) BudgetRepositoryInterface {
	log, ok := ctx.Value("logger").(*logger.Logger)
	if !ok {
		log = logger.NewLogger(false, logger.INFO)
	}
	return &BudgetRepository{
		db:  db,
		log: log,
	}
}

// Create saves a new budget in the database.
func (r *BudgetRepository) Create(ctx context.Context, b *entities.Budget) error {
	r.log.Log(logger.INFO, "/aws/demo-talent", "BudgetRepository", "Creating budget")
	fmt.Println(b.ID, b.Description, b.Amount, b.StartDate, b.EndDate)
	query := `
		INSERT INTO budgets (id, description, amount, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query, b.ID, b.Description, b.Amount, b.StartDate, b.EndDate)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", "Error creating budget")
		return fmt.Errorf("error creating budget: %w", err)
	}
	return nil
}

// GetByID retrieves a budget from the database by its ID.
func (r *BudgetRepository) GetByID(ctx context.Context, id string) (*entities.Budget, error) {
	query := `
		SELECT id, description, amount, start_date, end_date
		FROM budgets
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	b := &entities.Budget{}
	err := row.Scan(&b.ID, &b.Description, &b.Amount, &b.StartDate, &b.EndDate)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", "Error getting budget")
		return nil, fmt.Errorf("error getting budget: %w", err)
	}
	return b, nil
}

// Update updates a budget in the database.
func (r *BudgetRepository) Update(ctx context.Context, b *entities.Budget) error {
	r.log.Log(logger.INFO, "/aws/demo-talent", "BudgetRepository", "Updating budget")
	query := `
		UPDATE budgets
		SET description = $1, amount = $2, start_date = $3, end_date = $4
		WHERE id = $5
	`
	_, err := r.db.ExecContext(ctx, query, b.Description, b.Amount, b.StartDate, b.EndDate, b.ID)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", "Error updating budget")
		return fmt.Errorf("error updating budget: %w", err)
	}
	return nil
}

// Delete removes a budget from the database.
func (r *BudgetRepository) Delete(ctx context.Context, id string) error {
	r.log.Log(logger.INFO, "/aws/demo-talent", "BudgetRepository", "Deleting budget")
	query := `
		DELETE FROM budgets
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", "Error deleting budget")
		return fmt.Errorf("error deleting budget: %w", err)
	}
	return nil
}

// List retrieves a list of budgets from the database.
func (r *BudgetRepository) List(ctx context.Context, limit, offset int) ([]entities.Budget, error) {
	r.log.Log(logger.INFO, "/aws/demo-talent", "BudgetRepository", "Listing budgets")
	query := `
		SELECT id, description, amount, start_date, end_date
		FROM budgets
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", "Error listing budgets")
		return nil, fmt.Errorf("error listing budgets: %w", err)
	}
	defer rows.Close()

	var budgets []entities.Budget
	for rows.Next() {
		b := entities.Budget{}
		err := rows.Scan(&b.ID, &b.Description, &b.Amount, &b.StartDate, &b.EndDate)
		if err != nil {
			r.log.Log(logger.ERROR, "/aws/demo-talent", "BudgetRepository", "Error scanning budgets")
			return nil, fmt.Errorf("error scanning budgets: %w", err)
		}
		budgets = append(budgets, b)
	}
	return budgets, nil
}
