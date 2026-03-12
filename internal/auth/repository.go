package auth

import (
	"context"
	"soup/internal/store"
)

type Repository interface {
	FindByPhone(ctx context.Context, phone string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}

type repository struct {
	redis *store.Redis
	db    *store.Database
}

func NewRepository(redis *store.Redis, db *store.Database) Repository {
	return &repository{
		redis: redis,
		db:    db,
	}
}

func (r *repository) FindByPhone(ctx context.Context, phone string) (*User, error) {
	query := `
		SELECT id, phone, password, email, name, address, photo_url, is_admin, push_token, created_at
		FROM users 
		WHERE phone = $1
	`
	var user User
	err := r.db.DB.QueryRowContext(ctx, query, phone).Scan(
		&user.ID,
		&user.Phone,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Address,
		&user.PhotoURL,
		&user.IsAdmin,
		&user.PushToken,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(ctx context.Context, user User) (*User, error) {
	existingUser, err := r.FindByPhone(ctx, user.Phone)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}
	query := `
		INSERT INTO users (phone, password)
		VALUES ($1, $2)
		RETURNING id, phone, password,	email, name, address, photo_url, is_admin, push_token, created_at
	`
	var newUser User
	err = r.db.DB.QueryRowContext(ctx, query, user.Phone, user.Password).Scan(
		&newUser.ID,
		&newUser.Phone,
		&newUser.Password,
		&newUser.Email,
		&newUser.Name,
		&newUser.Address,
		&newUser.PhotoURL,
		&newUser.IsAdmin,
		&newUser.PushToken,
		&newUser.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
