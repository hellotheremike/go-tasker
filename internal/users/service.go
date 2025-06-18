package users

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]User, error)
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