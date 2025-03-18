package categories

import (
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (h *CategoryHandler) HandleCreateCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "No implementado aún",
	})
}

func (h *CategoryHandler) HandleGetCategories(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "No implementado aún",
	})
}

func (h *CategoryHandler) HandleUpdateCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "No implementado aún",
	})
}

func (h *CategoryHandler) HandleDeleteCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "No implementado aún",
	})
}
