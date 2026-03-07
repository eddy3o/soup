package auth

import (
	"context"
	"fmt"
	"soup/internal/store"

	"soup/internal/pkg/token"
	"soup/internal/pkg/utils"
)

type Service interface {
	Login(ctx context.Context, req UserLogin) (*token.Tokens, error)
	Logout(ctx context.Context, accessToken string, refreshToken string) error
	Register(ctx context.Context, phone string, hashedPassword string) (*User, error)
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
	user, err := s.repo.FindByPhone(ctx, req.Phone)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := utils.Verify(req.Password, user.Password); err != nil {
		fmt.Println("password mismatch")
		return nil, ErrInvalidCredentials
	}

	toks, err := token.IssueTokens(user.ID)

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

func (s *service) Register(ctx context.Context, phone string, hashedPassword string) (*User, error) {
	user, err := s.repo.Create(ctx, User{
		Phone:    phone,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
