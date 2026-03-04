package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soup/internal/pkg/token"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var req UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toks, err := h.service.Login(c, req)

	if err != nil {
		switch err {
		case ErrInvalidCredentials:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		}
	}
	token.SetAuthCookies(c, toks)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) Logout(c *gin.Context) {
	accessToken, _ := c.Cookie("access_token")
	refreshToken, _ := c.Cookie("refresh_token")

	h.service.Logout(c.Request.Context(), accessToken, refreshToken)

	token.ClearAuthCookies(c)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
