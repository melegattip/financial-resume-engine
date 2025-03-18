package transactions

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/models"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	db *gorm.DB
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

func (h *TransactionHandler) HandleCreateTransaction(c *gin.Context) {
	// Validar header x-caller-id
	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "x-caller-id header is required"})
		return
	}

	var request models.TransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar campos requeridos
	if request.TypeID == "" || request.Description == "" || request.Amount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type_id, description and amount are required"})
		return
	}

	// Validar tipo de transacción
	if request.TypeID != "income" && request.TypeID != "expense" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type_id must be either 'income' or 'expense'"})
		return
	}

	// Crear el modelo de transacción
	transactionModel := models.NewTransactionBuilder().
		SetTypeID(request.TypeID).
		SetDescription(request.Description).
		SetAmount(request.Amount).
		SetPayed(request.Payed).
		SetCategory(request.Category).
		Build()

	// Si hay fecha de expiración, intentar parsearla
	if request.ExpiryDate != "" {
		parsedTime, err := time.Parse(time.RFC3339, request.ExpiryDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiry_date format, must be RFC3339"})
			return
		}
		transactionModel.ExpiryDate = parsedTime
	}

	// Crear el servicio y ejecutar la operación
	service := NewCreateTransaction(NewTransactionRepository(h.db))
	response, err := service.Execute(userID, transactionModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *TransactionHandler) HandleListTransactions(c *gin.Context) {
	// Validar header x-caller-id
	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "x-caller-id header is required"})
		return
	}

	repo := NewTransactionRepository(h.db)
	transactions, err := repo.List(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list transactions"})
		return
	}

	if len(transactions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User has no transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
