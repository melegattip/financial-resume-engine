package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/usecases/incomes"
)

// IncomeHandler maneja las peticiones relacionadas con ingresos
type IncomeHandler struct {
	service incomes.IncomeService
}

func NewIncomeHandler(service incomes.IncomeService) *IncomeHandler {
	return &IncomeHandler{
		service: service,
	}
}

// CreateIncome godoc
// @Summary Crear un nuevo ingreso
// @Description Crea un nuevo ingreso para el usuario
// @Tags incomes
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param income body incomes.CreateIncomeRequest true "Datos del ingreso"
// @Success 201 {object} incomes.CreateIncomeResponse
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/incomes [post]
func (h *IncomeHandler) CreateIncome(c *gin.Context) {
	var request incomes.CreateIncomeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, apiErrors.NewBadRequest(err.Error()))
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, apiErrors.NewUnauthorizedRequest("User ID is required"))
		return
	}
	request.UserID = userID

	response, err := h.service.CreateIncome(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetIncome godoc
// @Summary Obtener un ingreso específico
// @Description Obtiene un ingreso por su ID
// @Tags incomes
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID del ingreso"
// @Success 200 {object} incomes.GetIncomeResponse
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/incomes/{id} [get]
func (h *IncomeHandler) GetIncome(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, apiErrors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	incomeID := c.Param("id")
	if incomeID == "" {
		c.JSON(http.StatusBadRequest, apiErrors.NewBadRequest("Income ID is required"))
		return
	}

	response, err := h.service.GetIncome(c.Request.Context(), userID, incomeID)
	if err != nil {
		c.JSON(http.StatusNotFound, apiErrors.NewResourceNotFound("Income not found"))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListIncomes godoc
// @Summary Listar ingresos
// @Description Obtiene una lista de todos los ingresos del usuario
// @Tags incomes
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Success 200 {object} incomes.ListIncomesResponse
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/incomes [get]
func (h *IncomeHandler) ListIncomes(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, apiErrors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	response, err := h.service.ListIncomes(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateIncome godoc
// @Summary Actualizar un ingreso
// @Description Actualiza los datos de un ingreso existente
// @Tags incomes
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID del ingreso"
// @Param income body incomes.UpdateIncomeRequest true "Datos actualizados del ingreso"
// @Success 200 {object} incomes.UpdateIncomeResponse
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/incomes/{id} [patch]
func (h *IncomeHandler) UpdateIncome(c *gin.Context) {
	var request incomes.UpdateIncomeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, apiErrors.NewBadRequest(err.Error()))
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, apiErrors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	incomeID := c.Param("id")
	if incomeID == "" {
		c.JSON(http.StatusBadRequest, apiErrors.NewBadRequest("Income ID is required"))
		return
	}

	response, err := h.service.UpdateIncome(c.Request.Context(), userID, incomeID, &request)
	if err != nil {
		c.JSON(http.StatusNotFound, apiErrors.NewResourceNotFound("Income not found"))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteIncome godoc
// @Summary Eliminar un ingreso
// @Description Elimina un ingreso existente
// @Tags incomes
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID del ingreso"
// @Success 204 "No Content"
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/incomes/{id} [delete]
func (h *IncomeHandler) DeleteIncome(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, apiErrors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	incomeID := c.Param("id")
	if incomeID == "" {
		c.JSON(http.StatusBadRequest, apiErrors.NewBadRequest("Income ID is required"))
		return
	}

	err := h.service.DeleteIncome(c.Request.Context(), userID, incomeID)
	if err != nil {
		c.JSON(http.StatusNotFound, apiErrors.NewResourceNotFound("Income not found"))
		return
	}

	c.Status(http.StatusNoContent)
}
