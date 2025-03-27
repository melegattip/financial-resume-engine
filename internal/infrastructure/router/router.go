package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melegattip/financial-resume-engine/internal/handlers"
)

// SetupRouter configura todas las rutas de la aplicación
func SetupRouter(incomeHandler *handlers.IncomeHandler, expenseHandler *handlers.ExpenseHandler, categoryHandler *handlers.CategoryHandler) *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Middleware para obtener el user_id del contexto
	router.Use(func(c *gin.Context) {
		// TODO: Implementar autenticación real
		c.Set("user_id", "test-user")
		c.Next()
	})

	// Grupo de rutas para la API v1
	v1 := router.Group("/api/v1")
	{
		// Rutas de categorías
		categories := v1.Group("/categories")
		{
			categories.POST("", categoryHandler.CreateCategory)
			categories.GET("", categoryHandler.GetCategories)
			categories.PATCH("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		// Rutas de ingresos
		incomes := v1.Group("/incomes")
		{
			incomes.POST("", incomeHandler.CreateIncome)
			incomes.GET("", incomeHandler.ListIncomes)
			incomes.GET("/:id", incomeHandler.GetIncome)
			incomes.PATCH("/:id", incomeHandler.UpdateIncome)
		}

		// Rutas de gastos
		expenses := v1.Group("/expenses")
		{
			expenses.POST("", expenseHandler.CreateExpense)
			expenses.GET("", expenseHandler.ListExpenses)
			expenses.GET("/unpaid", expenseHandler.ListUnpaidExpenses)
			expenses.GET("/by-due-date", expenseHandler.ListExpensesByDueDate)
			expenses.GET("/:id", expenseHandler.GetExpense)
			expenses.PATCH("/:id", expenseHandler.UpdateExpense)
			expenses.PATCH("/:id/paid", expenseHandler.MarkAsPaid)
		}
	}

	return router
}
