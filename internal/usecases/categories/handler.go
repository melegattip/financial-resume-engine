package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/repository"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

// CreateCategory godoc
// @Summary Crear una nueva categoría
// @Description Crea una nueva categoría con los datos proporcionados
// @Tags categories
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param category body CreateCategoryRequest true "Datos de la categoría"
// @Success 201 {object} domain.Category
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/categories [post]
func (h *CategoryHandler) HandleCreateCategory(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest(err.Error()))
		return
	}

	// Crear el modelo de categoría
	categoryModel := NewCategoryBuilder().
		SetName(request.Name).
		SetDescription(request.Description).
		Build()

	// Convertir a domain.Category
	domainCategory := &domain.Category{
		ID:          categoryModel.ID,
		Name:        categoryModel.Name,
		Description: categoryModel.Description,
		CreatedAt:   categoryModel.CreatedAt,
		UpdatedAt:   categoryModel.UpdatedAt,
	}

	// Crear el servicio y ejecutar la operación
	service := NewCreateCategory(repository.NewCategoryRepository(h.db))
	category, err := service.Execute(domainCategory)
	if err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorCreatingCategory.GetMessage(), logs.Tags{
			"name": request.Name,
		})
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategories godoc
// @Summary Obtener todas las categorías
// @Description Retorna una lista de todas las categorías disponibles
// @Tags categories
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Success 200 {array} domain.Category
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/categories [get]
func (h *CategoryHandler) HandleGetCategories(c *gin.Context) {
	service := NewListCategories(repository.NewCategoryRepository(h.db))
	categories, err := service.Execute()
	if err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorListingCategories.GetMessage(), logs.Tags{})
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary Actualizar una categoría
// @Description Actualiza una categoría existente con los datos proporcionados
// @Tags categories
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID de la categoría"
// @Param category body UpdateCategoryRequest true "Datos de actualización"
// @Success 200 {object} domain.Category
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/categories/{id} [patch]
func (h *CategoryHandler) HandleUpdateCategory(c *gin.Context) {
	id := c.Param("id")
	service := NewGetCategory(repository.NewCategoryRepository(h.db))
	category, err := service.Execute(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		logger.Error(c.Request.Context(), err, logs.ErrorListingCategories.GetMessage(), logs.Tags{"id": id})
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	var request struct {
		Name        *string `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest(err.Error()))
		return
	}

	if request.Name != nil {
		category.Name = *request.Name
	}
	if request.Description != nil {
		category.Description = *request.Description
	}

	updateService := NewUpdateCategory(repository.NewCategoryRepository(h.db))
	if err := updateService.Execute(category); err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorUpdatingCategory.GetMessage(), logs.Tags{
			"id":   id,
			"name": request.Name,
		})
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Eliminar una categoría
// @Description Elimina una categoría existente por su ID
// @Tags categories
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID de la categoría"
// @Success 204 "No Content"
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/categories/{id} [delete]
func (h *CategoryHandler) HandleDeleteCategory(c *gin.Context) {
	id := c.Param("id")
	service := NewDeleteCategory(repository.NewCategoryRepository(h.db))
	if err := service.Execute(id); err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorDeletingCategory.GetMessage(), logs.Tags{"id": id})
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
