package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler, authMiddleware gin.HandlerFunc) {
	r.GET("/me", authMiddleware, h.GetUserByID)
	r.PATCH("/me", authMiddleware, h.PatchUser)
}
