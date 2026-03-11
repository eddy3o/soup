package products

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *Handler, authMiddleware gin.HandlerFunc) {

	r.GET("/", authMiddleware, h.FindAll)
	// r.GET("/:id", h.FindByID)
	// r.POST("/", h.Create)
	// r.PUT("/:id", h.Update)
	// r.DELETE("/:id", h.Delete)
}
