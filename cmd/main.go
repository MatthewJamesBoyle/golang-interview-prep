package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/matthewjamesboyle/golang-interview-prep/internal/user"
	"log"
	"net/http"
)

func main() {

	runMigrations()

	svc, err := user.NewService("admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	h := user.Handler{Svc: *svc}

	http.HandleFunc("/user", h.AddUser)

	log.Println("starting http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runMigrations() {
	// Database connection string
	dbURL := "postgres://admin:admin@localhost/test_repo?sslmode=disable"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new instance of the PostgreSQL driver for migrate
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Database migration complete.")
}
