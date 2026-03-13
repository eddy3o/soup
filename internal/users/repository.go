package users

import (
	"context"
	"soup/internal/auth"
	"soup/internal/store"
)

type Repository interface {
	Get(ctx context.Context, id string) (*auth.User, error)
	Put(ctx context.Context, user auth.User) (*auth.User, error)
}

type repository struct {
	db *store.Database
}

func NewRepository(db *store.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (*auth.User, error) {
	query := `
		SELECT id, phone, email, name, address, photo_url, is_admin, push_token, created_at
		FROM users 
		WHERE id = $1
	`
	var user auth.User
	err := r.db.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Phone,
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

func (r *repository) Put(ctx context.Context, user auth.User) (*auth.User, error) {
	query := `
		INSERT INTO users (id, phone, password, email, name, address, photo_url, is_admin, push_token)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			phone = EXCLUDED.phone,
			password = EXCLUDED.password,
			email = EXCLUDED.email,
			name = EXCLUDED.name,
			address = EXCLUDED.address,
			photo_url = EXCLUDED.photo_url,
			is_admin = EXCLUDED.is_admin,
			push_token = EXCLUDED.push_token
		RETURNING id, phone, password, email, name, address, photo_url, is_admin, push_token, created_at
	`

	var newUser auth.User
	err := r.db.DB.QueryRowContext(ctx, query,
		user.ID,
		user.Phone,
		user.Password,
		user.Email,
		user.Name,
		user.Address,
		user.PhotoURL,
		user.IsAdmin,
		user.PushToken,
	).Scan(
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
