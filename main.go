package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"context"
	"github.com/demo-talent/internal/handlers"
	"github.com/demo-talent/internal/repository"
	"github.com/demo-talent/internal/services"
	"github.com/demo-talent/internal/routes"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
	"github.com/demo-talent/logger"
) 

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
 
	// Set up the logger
	isDebug := os.Getenv("DEBUG") == "true"
    logLevel := logger.INFO
    if isDebug {
        logLevel = logger.DEBUG 
    }
    log := logger.NewLogger(isDebug, logLevel) // Default to INFO
	ctx := context.Background()
	ctx = context.WithValue(ctx, "logger", log)

	// Set up the database connection
	log.Log(logger.INFO, "/aws/demo-talent", "main", "Starting the server")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Log(logger.FATAL, "/aws/demo-talent", "main", "Error connecting to the database")
	}
	defer db.Close()

	// Run database migrations
	if err := runMigrations(db, log); err != nil {
		log.Log(logger.FATAL, "/aws/demo-talent", "main", "Error running migrations")
	}

	repo := repository.NewExpenseRepository(ctx, db)
	svc := services.NewExpenseService(ctx, repo)
	expensesHandlers := handlers.NewExpensesRouter(ctx, svc) 
	

	// Set up the routes
	r := mux.NewRouter()
	routes.SetupExpensesRoutes(r, expensesHandlers)
	routes.SetupSwaggerRoutes(r)


	log.Log(logger.INFO, "/aws/demo-talent", "main", "Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Log(logger.FATAL, "/aws/demo-talent", "main", "Error starting the server")
	}
}

func runMigrations(db *sql.DB, log *logger.Logger) error {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
		message := fmt.Sprintf("Error creating migration driver: %v", err)
		log.Log(logger.ERROR, "/aws/demo-talent", "main", message)
        return fmt.Errorf("error creating migration driver: %w", err)
    }

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		message := fmt.Sprintf("Error creating migration instance: %v", err)
		log.Log(logger.ERROR, "/aws/demo-talent", "main", message)
		return fmt.Errorf("error creating migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		message := fmt.Sprintf("Error running migrations: %v", err)
		log.Log(logger.ERROR, "/aws/demo-talent", "main", message)
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
