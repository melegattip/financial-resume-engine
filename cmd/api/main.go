package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/melegattip/financial-resume-engine/docs"
	"github.com/melegattip/financial-resume-engine/internal/handlers"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/repository"
	"github.com/melegattip/financial-resume-engine/internal/usecases/expenses"
	"github.com/melegattip/financial-resume-engine/internal/usecases/incomes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Financial Resume Engine API
// @version 1.0
// @description API para gestionar ingresos y gastos personales
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	// Obtener variables de entorno
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "financial_resume")
	dbPort := getEnv("DB_PORT", "5432")

	// Construir DSN para la conexión a la base de datos
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Inicializar la conexión a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Inicializar repositorios
	incomeRepo := repository.NewIncomeRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)

	// Inicializar servicios
	incomeService := incomes.NewIncomeService(incomeRepo)
	expenseService := expenses.NewExpenseService(expenseRepo)

	// Inicializar handlers
	incomeHandler := handlers.NewIncomeHandler(incomeService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	// Configurar el router
	router := gin.Default()
	router.Use(gin.Recovery())

	// Configurar Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configurar rutas
	api := router.Group("/api/v1")
	{
		// Rutas de ingresos
		incomes := api.Group("/incomes")
		{
			incomes.POST("", incomeHandler.CreateIncome)
			incomes.GET("", incomeHandler.ListIncomes)
			incomes.GET("/:id", incomeHandler.GetIncome)
			incomes.PATCH("/:id", incomeHandler.UpdateIncome)
			incomes.DELETE("/:id", incomeHandler.DeleteIncome)
		}

		// Rutas de gastos
		expenses := api.Group("/expenses")
		{
			expenses.POST("", expenseHandler.CreateExpense)
			expenses.GET("", expenseHandler.ListExpenses)
			expenses.GET("/unpaid", expenseHandler.ListUnpaidExpenses)
			expenses.GET("/by-due-date", expenseHandler.ListExpensesByDueDate)
			expenses.GET("/:id", expenseHandler.GetExpense)
			expenses.PATCH("/:id", expenseHandler.UpdateExpense)
			expenses.PATCH("/:id/paid", expenseHandler.MarkAsPaid)
			expenses.DELETE("/:id", expenseHandler.DeleteExpense)
		}
	}

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv obtiene una variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
