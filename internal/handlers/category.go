package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	httputil "github.com/melegattip/financial-resume-engine/internal/infrastructure/http"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/repository"
	"github.com/melegattip/financial-resume-engine/internal/usecases/categories"
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
// @Param category body object{name=string} true "Datos de la categoría"
// @Success 201 {object} domain.Category
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var request struct {
		UserID string `json:"user_id" binding:"required"`
		Name   string `json:"name" binding:"required"`
	}

	userID, err := validateCallerID(c)
	if err != nil {
		httputil.HandleError(c, err)
		return
	}

	// Crear el modelo de categoría
	categoryModel := domain.NewCategoryBuilder().
		SetUserID(userID).
		SetName(request.Name).
		Build()

	// Crear el servicio y ejecutar la operación
	service := categories.NewCreateCategory(repository.NewCategoryRepository(h.db))
	category, err := service.Execute(categoryModel)
	if err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorCreatingCategory.GetMessage(), logs.Tags{
			"name": request.Name,
		})
		if errors.IsResourceAlreadyExists(err) {
			httputil.HandleError(c, errors.NewResourceAlreadyExists(err.Error()))
			return
		}
		httputil.HandleError(c, errors.NewInternalServerError(err.Error()))
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
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		httputil.HandleError(c, errors.NewUnauthorizedRequest("x-caller-id header is required"))
		return
	}

	service := categories.NewListCategories(repository.NewCategoryRepository(h.db))
	categories, err := service.Execute()
	if err != nil {
		if errors.IsResourceNotFound(err) {
			httputil.HandleError(c, errors.NewResourceNotFound(err.Error()))
			return
		}

		logger.Error(c.Request.Context(), err, logs.ErrorListingCategories.GetMessage(), logs.Tags{})
		httputil.HandleError(c, errors.NewInternalServerError(err.Error()))
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
// @Param category body object{name=string} true "Datos de actualización"
// @Success 200 {object} domain.Category
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/categories/{id} [patch]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	_, err := validateCallerID(c)
	if err != nil {
		httputil.HandleError(c, err)
		return
	}

	categoryName := c.GetHeader("x-category-name")
	if categoryName == "" {
		httputil.HandleError(c, errors.NewBadRequest("x-category-name header is required"))
		return
	}

	service := categories.NewGetCategory(repository.NewCategoryRepository(h.db))
	category, err := service.Execute(categoryName)
	if err != nil {
		if errors.IsResourceNotFound(err) {
			httputil.HandleError(c, errors.NewResourceNotFound("Category not found"))
			return
		}
		logger.Error(c.Request.Context(), err, logs.ErrorListingCategories.GetMessage(), logs.Tags{"name": categoryName})
		httputil.HandleError(c, errors.NewInternalServerError(err.Error()))
		return
	}

	var request struct {
		Name *string `json:"name,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		httputil.HandleError(c, errors.NewBadRequest(err.Error()))
		return
	}

	if request.Name != nil {
		category.Name = *request.Name
	}

	updateService := categories.NewUpdateCategory(repository.NewCategoryRepository(h.db))
	if err := updateService.Execute(category); err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorUpdatingCategory.GetMessage(), logs.Tags{
			"name": request.Name,
		})
		httputil.HandleError(c, errors.NewInternalServerError(err.Error()))
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
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	service := categories.NewDeleteCategory(repository.NewCategoryRepository(h.db))
	if err := service.Execute(id); err != nil {
		logger.Error(c.Request.Context(), err, logs.ErrorDeletingCategory.GetMessage(), logs.Tags{"id": id})
		httputil.HandleError(c, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func validateCallerID(c *gin.Context) (string, error) {
	var request struct {
		UserID string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		return "", errors.NewBadRequest(err.Error())
	}

	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		return "", errors.NewUnauthorizedRequest("x-caller-id header is required")
	}

	if request.UserID != userID {
		return "", errors.NewUnauthorizedRequest("user_id access denied")
	}

	return userID, nil
}
