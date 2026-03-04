package auth

import (
	"context"
	"soup/internal/store"
)

type Repository interface {
	FindByPhone(ctx context.Context, phone string) (*User, error)
}

type repository struct {
	redis *store.Redis
}

func NewRepository(redis *store.Redis) Repository {
	return &repository{
		redis: redis,
	}
}

func (r *repository) FindByPhone(ctx context.Context, phone string) (*User, error) {
	if phone == "4436404393" {
		return &User{
			ID:       "1",
			Phone:    "4436404393",
			Password: "password",
		}, nil
	}
	return nil, nil
}

