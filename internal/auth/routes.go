package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)
}
