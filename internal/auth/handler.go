package auth

import (
	"net/http"
	"soup/internal/pkg/token"
	"soup/internal/pkg/utils"
	"soup/internal/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
	rds     *store.Redis
}

func NewHandler(service Service, redis *store.Redis) *Handler {
	return &Handler{
		service: service,
		rds:     redis,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, toks, err := h.service.Login(c, req)

	if err != nil {
		switch err {
		case ErrInvalidCredentials:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		case ErrUserNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			return
		}

	}
	token.SetAuthCookies(c, toks)
	token.Persist(c, h.rds, toks)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  toks.Access,
		"refresh_token": toks.Refresh,
		"user":          user,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	accessToken, _ := c.Cookie("access_token")
	refreshToken, _ := c.Cookie("refresh_token")

	h.service.Logout(c.Request.Context(), accessToken, refreshToken)

	token.ClearAuthCookies(c)
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Register(c *gin.Context) {
	var req UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.Hash(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	_, err = h.service.Register(c.Request.Context(), req.Phone, hashedPassword)
	if err != nil {
		switch err {
		case ErrUserAlreadyExists:
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
