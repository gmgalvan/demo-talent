package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/demo-talent/handlers"
	"github.com/demo-talent/repository"
	"github.com/demo-talent/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
	"github.com/demo-talent/logger"
)

func main() {
	isDebug := os.Getenv("DEBUG") == "true"
    logLevel := logger.INFO
    if isDebug {
        logLevel = logger.DEBUG 
    }
    log := logger.NewLogger(isDebug, logLevel) // Default to INFO

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "logger", log)

	log.Log(logger.INFO, "/aws/demo-talent", "main", "Starting the server")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Log(logger.FATAL, "/aws/demo-talent", "main", "Error connecting to the database")
	}
	defer db.Close()

	// Run database migrations
	if err := runMigrations(db); err != nil {
		log.Log(logger.FATAL, "/aws/demo-talent", "main", "Error running migrations")
	}

	repo := repository.NewExpenseRepository(ctx, db)
	svc := services.NewExpenseService(ctx, repo)

	r := mux.NewRouter() 

	// Register the expense handlers routes
	r.HandleFunc("/expenses", handlers.CreateExpense(svc)).Methods("POST") // with ID
	r.HandleFunc("/expenses", handlers.GetExpense(svc)).Methods("GET")
	r.HandleFunc("/expenses", handlers.UpdateExpense(svc)).Methods("PUT")
	r.HandleFunc("/expenses", handlers.DeleteExpense(svc)).Methods("DELETE")
	r.HandleFunc("/", handlers.HelloWorld).Methods("GET")

	opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	sh := middleware.Redoc(opts, nil)
	r.Handle("/docs", sh).Methods("GET")
	r.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	log.Log(logger.INFO, "/aws/demo-talent", "main", "Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Log(logger.FATAL, "/aws/demo-talent", "main", "Error starting the server")
	}
}

func runMigrations(db *sql.DB) error {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("error creating migration driver: %w", err)
    }

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
