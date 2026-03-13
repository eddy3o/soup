package users

import (
	"context"
	"soup/internal/auth"
)

type Service interface {
	GetUserByID(ctx context.Context, id string) (*auth.User, error)
	// Put(ctx context.Context, user auth.User) (*auth.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserByID(ctx context.Context, id string) (*auth.User, error) {
	return s.repo.Get(ctx, id)
}
