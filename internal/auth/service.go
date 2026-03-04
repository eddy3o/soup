package auth

import (
	"context"
	"errors"
	"soup/internal/store"

	"soup/internal/pkg/token"
)

var demoUser = User{ID: "1", Phone: "4436404393", Password: "password"}

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service interface {
	Login(ctx context.Context, req UserLogin) (*token.Tokens, error)
	Logout(ctx context.Context, accessToken string, refreshToken string) error
}

type service struct {
	repo  Repository
	redis *store.Redis
}

func NewService(repo Repository, redis *store.Redis) Service {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) Login(ctx context.Context, req UserLogin) (*token.Tokens, error) {
	if req.Phone != demoUser.Phone || req.Password != demoUser.Password {
		return nil, ErrInvalidCredentials
	}

	toks, err := token.IssueTokens(demoUser.ID)

	return toks, err
}

func (s *service) Logout(ctx context.Context, accessToken string, refreshToken string) error {
	if accessToken != "" {
		if claims, err := token.ParseAccess(accessToken); err == nil {
			_ = s.redis.DelJTI(ctx, "access:"+claims.ID)
		}
	}
	if refreshToken != "" {
		if claims, err := token.ParseRefresh(refreshToken); err == nil {
			_ = s.redis.DelJTI(ctx, "refresh:"+claims.ID)
		}
	}
	return nil
}
