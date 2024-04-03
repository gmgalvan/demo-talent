package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/demo-talent/internal/entities"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/demo-talent/logger"
)

type ExpenseRepositoryInterface interface {
	Create(ctx context.Context, e *entities.Expense) error
	GetByID(ctx context.Context, id string) (*entities.Expense, error)
	Update(ctx context.Context, e *entities.Expense) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]entities.Expense, error)
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
        INSERT INTO expenses (id, description, amount, date_creation, currency, name)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := r.db.ExecContext(ctx, query, e.ID, e.Description, e.Amount, e.DateCreation, e.Currency, e.Name)
	if err != nil {
		r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error creating expense")
		return fmt.Errorf("error creating expense: %w", err)
	}
	return nil
}

// GetByID retrieves an expense from the database by its ID.
func (r *ExpenseRepository) GetByID(ctx context.Context, id string) (*entities.Expense, error) {
	query := `
        SELECT id, description, amount, date_creation, currency, name
        FROM expenses
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	var e entities.Expense
	err := row.Scan(&e.ID, &e.Description, &e.Amount, &e.DateCreation, &e.Currency, &e.Name)
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
		SET description = $1, amount = $2, date_creation = $3, currency = $4, name = $5
		WHERE id = $6
    `
	_, err := r.db.ExecContext(ctx, query, e.Description, e.Amount, e.DateCreation, e.Currency, e.Name, e.ID)
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

// List retrieves expenses from the database with pagination.
func (r *ExpenseRepository) List(ctx context.Context, limit, offset int) ([]entities.Expense, error) {
    query := `
        SELECT id, description, amount, date_creation, currency, name
        FROM expenses
        ORDER BY date_creation DESC
        LIMIT $1 OFFSET $2
    `
    rows, err := r.db.QueryContext(ctx, query, limit, offset)
    if err != nil {
        r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error listing expenses")
        return nil, fmt.Errorf("error listing expenses: %w", err)
    }
    defer rows.Close()

    var expenses []entities.Expense
    for rows.Next() {
        var e entities.Expense
        err := rows.Scan(&e.ID, &e.Description, &e.Amount, &e.DateCreation, &e.Currency, &e.Name)
        if err != nil {
            r.log.Log(logger.ERROR, "/aws/demo-talent", "ExpenseRepository", "Error scanning expenses")
            return nil, fmt.Errorf("error scanning expenses: %w", err)
        }
        expenses = append(expenses, e)
    }

    return expenses, nil
}