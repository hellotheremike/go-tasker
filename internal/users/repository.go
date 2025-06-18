package users

import (
	"context"
	"database/sql"
	"log"

	"github.com/hellotheremike/go-tasker/internal/utils"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user UserRegister) (User, error)
	LoginUser(ctx context.Context, login UserLogin) (AuthResponse, error)
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
		err = rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repo) CreateUser(ctx context.Context, user UserRegister) (User, error) {
	var newUser User
	hashedPassword, err := utils.HashPassword(string(user.Password))
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return User{}, err
	}

	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, created_at`
	err = r.db.QueryRowContext(ctx, query, user.Email, hashedPassword).
		Scan(&newUser.ID, &newUser.CreatedAt)

	return newUser, err
}

func (r *repo) LoginUser(ctx context.Context, login UserLogin) (AuthResponse, error) {
	var user User

	query := `SELECT id, email, password, created_at FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, login.Email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Printf("Failed to query Login: %v", err)
		return AuthResponse{}, err
	}

	err = utils.CheckPassword(user.Password, login.Password)
	if err != nil {
		log.Printf("wrong password: %v", err)
		return AuthResponse{}, err
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return AuthResponse{}, err
	}

	return AuthResponse{Token: token}, nil
}
