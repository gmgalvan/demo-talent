package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/demo-talent/entities"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/demo-talent/logger"
)

type ExpenseRepositoryInterface interface {
	Create(ctx context.Context, e *entities.Expense) error
	GetByID(ctx context.Context, id string) (*entities.Expense, error)
	Update(ctx context.Context, e *entities.Expense) error
	Delete(ctx context.Context, id string) error
}

type ExpenseRepository struct {
	db *sql.DB
	log  *logger.Logger
}

// NewExpenseRepository creates a new instance of ExpenseRepository.
func NewExpenseRepository(ctx context.Context, db *sql.DB) ExpenseRepositoryInterface {
	log, ok := ctx.Value("logger").(*logger.Logger)
    if !ok {
		log = logger.NewLogger(false, logger.INFO)
    }
	return &ExpenseRepository{
		db: db,
		log: log,
	}
}

// Create saves a new expense in the database.
func (r *ExpenseRepository) Create(ctx context.Context, e *entities.Expense) error {
	r.log.Log(logger.INFO, "/aws/demo-talent", "ExpenseRepository", "Creating expense")
	query := `
        INSERT INTO expenses (id, description, amount, date_creation)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.ExecContext(ctx, query, e.ID, e.Description, e.Amount, e.DateCreation)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error creating expense")
		return fmt.Errorf("error creating expense: %w", err)
	}
	return nil
}

// GetByID retrieves an expense from the database by its ID.
func (r *ExpenseRepository) GetByID(ctx context.Context, id string) (*entities.Expense, error) {
	query := `
        SELECT id, description, amount, date_creation
        FROM expenses
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	var e entities.Expense
	err := row.Scan(&e.ID, &e.Description, &e.Amount, &e.DateCreation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("expense not found with ID: %s", id)
		}
		r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error retrieving expense")
		return nil, fmt.Errorf("error retrieving expense: %w", err)
	}

	return &e, nil
}

// Update updates an existing expense in the database.
func (r *ExpenseRepository) Update(ctx context.Context, e *entities.Expense) error {
	query := `
        UPDATE expenses
        SET description = $1, amount = $2
        WHERE id = $3
    `
	_, err := r.db.ExecContext(ctx, query, e.Description, e.Amount, e.ID)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error updating expense")
		return fmt.Errorf("error updating expense: %w", err)
	}
	return nil
}

// Delete removes an expense from the database by its ID.
func (r *ExpenseRepository) Delete(ctx context.Context, id string) error {
	query := `
        DELETE FROM expenses
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error deleting expense")
		return fmt.Errorf("error deleting expense: %w", err)
	}
	return nil
}
