package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Implementar lógica de creación
	c.JSON(http.StatusCreated, category)
}

func ListCategories(c *gin.Context) {
	// TODO: Implementar lógica de listado
	categories := []Category{}
	c.JSON(http.StatusOK, categories)
}
