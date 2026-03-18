package orders

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler, authMiddleware gin.HandlerFunc) {
	r.POST("", authMiddleware, h.CreateOrder)
}
