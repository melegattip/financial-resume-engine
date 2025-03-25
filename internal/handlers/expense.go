package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/usecases/expenses"
)

// ExpenseHandler maneja las peticiones relacionadas con gastos
type ExpenseHandler struct {
	service expenses.ExpenseService
}

func NewExpenseHandler(service expenses.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

// CreateExpense godoc
// @Summary Crear un nuevo gasto
// @Description Crea un nuevo gasto para el usuario
// @Tags expenses
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param expense body expenses.CreateExpenseRequest true "Datos del gasto"
// @Success 201 {object} expenses.CreateExpenseResponse
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/expenses [post]
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var request expenses.CreateExpenseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest(err.Error()))
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedRequest("User ID is required"))
		return
	}
	request.UserID = userID

	response, err := h.service.CreateExpense(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetExpense godoc
// @Summary Obtener un gasto específico
// @Description Obtiene un gasto por su ID
// @Tags expenses
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID del gasto"
// @Success 200 {object} expenses.GetExpenseResponse
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/expenses/{id} [get]
func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	expenseID := c.Param("id")
	if expenseID == "" {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest("Expense ID is required"))
		return
	}

	response, err := h.service.GetExpense(c.Request.Context(), userID, expenseID)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.NewResourceNotFound("Expense not found"))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListExpenses godoc
// @Summary Listar gastos
// @Description Obtiene una lista de todos los gastos del usuario
// @Tags expenses
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Success 200 {object} expenses.ListExpensesResponse
// @Failure 401 {object} errors.UnauthorizedRequest
// @Router /api/v1/expenses [get]
func (h *ExpenseHandler) ListExpenses(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	response, err := h.service.ListExpenses(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ExpenseHandler) ListUnpaidExpenses(c *gin.Context) {
	userID := c.GetString("user_id")

	response, err := h.service.ListUnpaidExpenses(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ExpenseHandler) ListExpensesByDueDate(c *gin.Context) {
	userID := c.GetString("user_id")

	response, err := h.service.ListExpensesByDueDate(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateExpense godoc
// @Summary Actualizar un gasto
// @Description Actualiza los datos de un gasto existente
// @Tags expenses
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID del gasto"
// @Param expense body expenses.UpdateExpenseRequest true "Datos actualizados del gasto"
// @Success 200 {object} expenses.UpdateExpenseResponse
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/expenses/{id} [patch]
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	var request expenses.UpdateExpenseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest(err.Error()))
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	expenseID := c.Param("id")
	if expenseID == "" {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest("Expense ID is required"))
		return
	}

	response, err := h.service.UpdateExpense(c.Request.Context(), userID, expenseID, &request)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.NewResourceNotFound("Expense not found"))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ExpenseHandler) MarkAsPaid(c *gin.Context) {
	userID := c.GetString("user_id")
	expenseID := c.Param("id")

	response, err := h.service.MarkAsPaid(c.Request.Context(), userID, expenseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteExpense godoc
// @Summary Eliminar un gasto
// @Description Elimina un gasto existente
// @Tags expenses
// @Accept json
// @Produce json
// @Param x-caller-id header string true "ID del usuario"
// @Param id path string true "ID del gasto"
// @Success 204 "No Content"
// @Failure 400 {object} errors.BadRequest
// @Failure 401 {object} errors.UnauthorizedRequest
// @Failure 404 {object} errors.ResourceNotFound
// @Router /api/v1/expenses/{id} [delete]
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedRequest("User ID is required"))
		return
	}

	expenseID := c.Param("id")
	if expenseID == "" {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest("Expense ID is required"))
		return
	}

	err := h.service.DeleteExpense(c.Request.Context(), userID, expenseID)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.NewResourceNotFound("Expense not found"))
		return
	}

	c.Status(http.StatusNoContent)
}
