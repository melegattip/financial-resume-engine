package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

func (h *CategoryHandler) HandleCreateCategory(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear el modelo de categoría
	categoryModel := NewCategoryBuilder().
		SetName(request.Name).
		SetDescription(request.Description).
		Build()

	// Crear el servicio y ejecutar la operación
	service := NewCreateCategory(NewCategoryRepository(h.db))
	category, err := service.Execute(categoryModel)
	if err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorCreatingCategory.GetMessage(), logs.Tags{
			"name": request.Name,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) HandleGetCategories(c *gin.Context) {
	service := NewListCategories(NewCategoryRepository(h.db))
	categories, err := service.Execute()
	if err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorListingCategories.GetMessage(), logs.Tags{})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) HandleUpdateCategory(c *gin.Context) {
	id := c.Param("id")
	service := NewGetCategory(NewCategoryRepository(h.db))
	category, err := service.Execute(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		logger.Error(c.Request.Context(), err, logs.ErrorListingCategories.GetMessage(), logs.Tags{"id": id})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting category"})
		return
	}

	var request struct {
		Name        *string `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name != nil {
		category.Name = *request.Name
	}
	if request.Description != nil {
		category.Description = *request.Description
	}

	updateService := NewUpdateCategory(NewCategoryRepository(h.db))
	if err := updateService.Execute(category); err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorUpdatingCategory.GetMessage(), logs.Tags{
			"id":   id,
			"name": request.Name,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) HandleDeleteCategory(c *gin.Context) {
	id := c.Param("id")
	service := NewDeleteCategory(NewCategoryRepository(h.db))
	if err := service.Execute(id); err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorDeletingCategory.GetMessage(), logs.Tags{"id": id})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
