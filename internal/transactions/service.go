package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Implementar lógica de creación
	c.JSON(http.StatusCreated, transaction)
}

func ListTransactions(c *gin.Context) {
	// TODO: Implementar lógica de listado
	transactions := []Transaction{}
	c.JSON(http.StatusOK, transactions)
}

func UpdateTransaction(c *gin.Context) {
	transactionID := c.Param("transaction_id")
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Implementar lógica de actualización
	c.JSON(http.StatusOK, gin.H{"id": transactionID, "message": "Transaction updated"})
}
