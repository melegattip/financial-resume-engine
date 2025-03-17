// cmd/app/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/categories"
	"github.com/melegattip/financial-resume-engine/internal/reports"
	"github.com/melegattip/financial-resume-engine/internal/transactions"
)

func main() {
	router := gin.Default()

	// Rutas para transacciones
	router.POST("/api/v1/transactions", transactions.CreateTransaction)
	router.GET("/api/v1/transactions", transactions.HandleListTransactions)
	router.PATCH("/api/v1/transactions/:transaction_id", transactions.HandleUpdateTransaction)

	// Rutas para categorías
	router.POST("/api/v1/categories", categories.CreateCategory)
	router.GET("/api/v1/categories", categories.ListCategories)

	// Rutas para reportes
	router.GET("/api/v1/reports/financial", reports.GenerateFinancialReport)

	router.Run(":8080")
}
