package orders

import "github.com/gin-gonic/gin"

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req OrderCreateRequest
	userID := c.GetString("userID")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := h.Service.CreateOrder(userID, req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, order)
}
