package users

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user UserRegister) (User, error)
	LoginUser(ctx context.Context, login UserLogin) (AuthResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) CreateUser(ctx context.Context, user UserRegister) (User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *service) LoginUser(ctx context.Context, login UserLogin) (AuthResponse, error) {
	return s.repo.LoginUser(ctx, login)
}
