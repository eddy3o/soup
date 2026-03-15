package users

import (
	"context"
	"soup/internal/auth"
	"soup/internal/store"
)

type Repository interface {
	Get(ctx context.Context, id string) (*auth.User, error)
	Patch(ctx context.Context, user auth.User) (*auth.User, error)
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

func (r *repository) Patch(ctx context.Context, user auth.User) (*auth.User, error) {
	query := `
	  UPDATE users
	  SET email = COALESCE(NULLIF($2, ''), email),
	      name = COALESCE(NULLIF($3, ''), name),
	      address = COALESCE(NULLIF($4, ''), address),
	      photo_url = COALESCE(NULLIF($5, ''), photo_url)
	  WHERE id = $1
	  RETURNING id, phone, password, email, name, address, photo_url, is_admin, push_token, created_at
	`

	var newUser auth.User
	err := r.db.DB.QueryRowContext(ctx, query,
		user.ID,
		user.Email.String,
		user.Name.String,
		user.Address.String,
		user.PhotoURL.String,
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
