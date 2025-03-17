package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/categories"
)

func CreateTransaction(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Implementar lógica de creación
	err := categories.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}
