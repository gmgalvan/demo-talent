package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/demo-talent/entities"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type ExpenseRepositoryInterface interface {
	Create(ctx context.Context, e *entities.Expense) error
	GetByID(ctx context.Context, id string) (*entities.Expense, error)
	Update(ctx context.Context, e *entities.Expense) error
	Delete(ctx context.Context, id string) error
}

type ExpenseRepository struct {
	db *sql.DB
}

// NewExpenseRepository creates a new instance of ExpenseRepository.
func NewExpenseRepository(db *sql.DB) ExpenseRepositoryInterface {
	return &ExpenseRepository{db: db}
}

// Create saves a new expense in the database.
func (r *ExpenseRepository) Create(ctx context.Context, e *entities.Expense) error {
	query := `
        INSERT INTO expenses (id, description, amount, date_creation)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.ExecContext(ctx, query, e.ID, e.Description, e.Amount, e.DateCreation)
	if err != nil {
		log.Printf("Error creating expense: %v", err)
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
		log.Printf("Error retrieving expense: %v", err)
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
		log.Printf("Error updating expense: %v", err)
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
		log.Printf("Error deleting expense: %v", err)
		return fmt.Errorf("error deleting expense: %w", err)
	}
	return nil
}
