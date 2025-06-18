package users

import (
	"context"
	"database/sql"
	"log"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{db: db}
}


func (r *repo) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT usename FROM pg_catalog.pg_user;`)
	if err != nil {
		log.Fatal("Failed to query GetUsers:", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}