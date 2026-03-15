package users

import (
	"database/sql"
	"errors"
	"net/http"
	"soup/internal/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.service.GetUserByID(c.Request.Context(), id.(string))

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) PatchUser(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var userUpdate UserUpdate
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user := auth.User{
		ID: id.(string),
		Name: sql.NullString{
			String: userUpdate.Name,
			Valid:  userUpdate.Name != "",
		},
		Address: sql.NullString{
			String: userUpdate.Address,
			Valid:  userUpdate.Address != "",
		},
		Email: sql.NullString{
			String: userUpdate.Email,
			Valid:  userUpdate.Email != "",
		},
		PhotoURL: sql.NullString{
			String: userUpdate.PhotoURL,
			Valid:  userUpdate.PhotoURL != "",
		},
	}

	updatedUser, err := h.service.PatchUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
