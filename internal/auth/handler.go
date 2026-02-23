package auth

import (
	"net/http"
	"soup/internal/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Redis *store.Redis
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var demoUser = User{ID: "1", Username: "user", Password: "pass"}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewHandler(redis *store.Redis) *Handler {
	return &Handler{Redis: redis}
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
