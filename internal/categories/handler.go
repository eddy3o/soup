package categories

import "github.com/gin-gonic/gin"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := (h.service).GetCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get categories"})
		return
	}
	c.JSON(200, categories)
}
