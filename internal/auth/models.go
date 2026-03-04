package auth

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	PhotoURL  string    `json:"photo_url"`
	IsAdmin   string    `json:"is_admin"`
	PushToken string    `json:"push_token"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogin struct {
	Phone    string `json:"phone" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *UserRegister) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}
	return nil
}
