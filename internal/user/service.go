package user

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type service struct {
	dbUser     string
	dbPassword string
}

func NewService(dbUser, dbPassword string) (*service, error) {
	if dbUser == "" {
		return nil, errors.New("dbUser was empty")
	}
	return &service{dbUser: dbUser, dbPassword: dbPassword}, nil
}

type User struct {
	Name     string
	Password string
}

func (s *service) AddUser(u User) (string, error) {
	// SMELL: direct connection to DB, consider ORM
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost/test_repo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id string
	// SMELL: direct SQL, consider ORM CRUD
	// SMELL: password string is saved directly, consider encryption
	// SMELL: in 000001_create_users_table.up.sql username should be made unique and validated
	q := "INSERT INTO users (username, password) VALUES ('" + u.Name + "', '" + u.Password + "') RETURNING id"

	err = db.QueryRow(q).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
