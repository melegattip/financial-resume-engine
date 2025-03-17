package main

import (
    "github.com/gin-gonic/gin"
    "github.com/tu_usuario/tu_proyecto/internal/transactions"
)

func main() {
    router := gin.Default()
    router.POST("/api/transactions", transactions.CreateTransaction)
    router.Run(":8080")
}
