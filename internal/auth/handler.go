package auth

import (
	"context"
	"net/http"
	"soup/internal/pkg/token"
	"soup/internal/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Redis *store.Redis
}

var demoUser = User{ID: "1", Username: "user", PasswordHash: "pass"}

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

	if req.Username != demoUser.Username || req.Password != demoUser.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	toks, err := token.IssueTokens(demoUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not issue tokens"})
		return
	}

	if err = token.Persist(c.Request.Context(), h.Redis, toks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not persist tokens"})
		return
	}

	token.SetAuthCookies(c, toks)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *Handler) Logout(c *gin.Context) {
	acc, _ := c.Cookie("access_token")
	ref, _ := c.Cookie("refresh_token")
	ctx := context.Background()

	if acc != "" {
		if claims, err := token.ParseAccess(acc); err == nil {
			_ = h.Redis.DelJTI(ctx, "access:"+claims.ID)
		}
	}
	if ref != "" {
		if claims, err := token.ParseRefresh(ref); err == nil {
			_ = h.Redis.DelJTI(ctx, "refresh:"+claims.ID)
		}
	}
	token.ClearAuthCookies(c)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
