package auth

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID        string         `json:"id"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	Name      sql.NullString `json:"name"`
	Address   sql.NullString `json:"address"`
	Email     sql.NullString `json:"email"`
	PhotoURL  sql.NullString `json:"photo_url"`
	IsAdmin   string         `json:"is_admin"`
	PushToken sql.NullString `json:"push_token"`
	CreatedAt time.Time      `json:"created_at"`
}

type UserLogin struct {
	Phone    string `json:"phone" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRegister struct {
	Phone    string `json:"phone" binding:"required,min=10"`
	Password string `json:"password" binding:"required,min=6"`
}

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrFailedToHashPassword = errors.New("failed to hash password")
)
