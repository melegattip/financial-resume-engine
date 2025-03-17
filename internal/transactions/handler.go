package transactions

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "Transaction created"})
}
