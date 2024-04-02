package main

import (
	"database/sql"
	"fmt"
	"log"
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
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Run database migrations
	if err := runMigrations(db); err != nil {
		log.Println("Database connection variables:")
        log.Println("DB_HOST:", dbHost)
        log.Println("DB_PORT:", dbPort)
        log.Println("DB_USER:", dbUser)
        log.Println("DB_PASSWORD:", dbPassword)
        log.Println("DB_NAME:", dbName)
        log.Fatal("Error running migrations:", err)
	}

	repo := repository.NewExpenseRepository(db)
	svc := services.NewExpenseService(repo)

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

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting the server:", err)
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
