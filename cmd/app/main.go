// cmd/app/main.go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/melegattip/financial-resume-engine/internal/config"
	"github.com/melegattip/financial-resume-engine/internal/transactions"
)

func main() {
	router := gin.Default()

	// Inicializar la base de datos
	db := config.InitDB()

	// Inicializar handlers
	transactionHandler := transactions.NewTransactionHandler(db)

	// Rutas para transacciones
	router.POST("/api/v1/transactions", transactionHandler.HandleCreateTransaction)
	router.GET("/api/v1/transactions", transactionHandler.HandleListTransactions)
	//router.PATCH("/api/v1/transactions/:transaction_id", transactionHandler.HandleUpdateTransaction)

	// Rutas para categorías
	//router.POST("/api/v1/categories", categories.CreateCategory)
	//router.GET("/api/v1/categories", categories.ListCategories)

	// Rutas para reportes
	//router.GET("/api/v1/reports/financial", reports.GenerateFinancialReport)

	router.Run(":8080")
}
