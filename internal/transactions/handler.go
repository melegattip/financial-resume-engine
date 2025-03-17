package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateTransaction(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created"})
}

func HandleListTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Transactions listed"})
}


