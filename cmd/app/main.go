// cmd/app/main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/melegattip/financial-resume-engine/internal/config"
	httpErrors "github.com/melegattip/financial-resume-engine/internal/infrastructure/http"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
	"github.com/melegattip/financial-resume-engine/internal/usecases/categories"
	"github.com/melegattip/financial-resume-engine/internal/usecases/reports"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/create"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/delete"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/get"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/list"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/update"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inicializar el logger
	logger.Init()

	// Inicializar la base de datos
	db := config.InitDB()

	// Crear el router
	router := gin.Default()

	// Inicializar handlers
	transactionRepo := create.NewTransactionRepository(db)
	createTransactionHandler := &create.CreateTransaction{
		Service: &create.CreateImp{
			TransactionRepository: transactionRepo,
		},
	}
	listTransactionHandler := &list.ListTransactions{
		Service: &list.ListImp{
			TransactionRepository: list.NewTransactionRepository(db),
		},
	}
	getTransactionHandler := &get.GetTransaction{
		Service: &get.GetImp{
			TransactionRepository: get.NewTransactionRepository(db),
		},
	}
	updateTransactionHandler := &update.UpdateTransaction{
		Service: &update.UpdateImp{
			TransactionRepository: update.NewTransactionRepository(db),
		},
	}
	deleteTransactionHandler := &delete.DeleteTransaction{
		Service: &delete.DeleteImp{
			TransactionRepository: delete.NewTransactionRepository(db),
		},
	}
	categoryHandler := categories.NewCategoryHandler(db)
	reportHandler := reports.NewReportHandler(db)

	// Rutas para transacciones
	router.POST("/api/v1/transactions", func(c *gin.Context) {
		if err := createTransactionHandler.Handle(c); err != nil {
			httpErrors.HandleError(c, err)
		}
	})

	router.GET("/api/v1/transactions", func(c *gin.Context) {
		if err := listTransactionHandler.Handle(c); err != nil {
			httpErrors.HandleError(c, err)
		}
	})

	router.GET("/api/v1/transactions/:id", func(c *gin.Context) {
		if err := getTransactionHandler.Handle(c); err != nil {
			httpErrors.HandleError(c, err)
		}
	})

	router.PATCH("/api/v1/transactions/:id", func(c *gin.Context) {
		if err := updateTransactionHandler.Handle(c); err != nil {
			httpErrors.HandleError(c, err)
		}
	})

	router.DELETE("/api/v1/transactions/:id", func(c *gin.Context) {
		if err := deleteTransactionHandler.Handle(c); err != nil {
			httpErrors.HandleError(c, err)
		}
	})

	// Rutas para categorías
	router.POST("/api/v1/categories", categoryHandler.HandleCreateCategory)
	router.GET("/api/v1/categories", categoryHandler.HandleGetCategories)
	router.PATCH("/api/v1/categories/:id", categoryHandler.HandleUpdateCategory)
	router.DELETE("/api/v1/categories/:id", categoryHandler.HandleDeleteCategory)

	// Rutas para reportes
	router.GET("/api/v1/reports/financial", reportHandler.HandleGenerateReport)

	// Iniciar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		logger.Error(nil, err, "Error starting server", logger.Tags{"error": err.Error()})
		os.Exit(1)
	}
}
