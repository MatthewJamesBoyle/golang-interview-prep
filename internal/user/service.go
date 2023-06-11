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
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost/test_repo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id string
	q := "INSERT INTO users (username, password) VALUES ('" + u.Name + "', '" + u.Password + "') RETURNING id"

	err = db.QueryRow(q).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
