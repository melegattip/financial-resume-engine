// cmd/app/main.go
package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/melegattip/financial-resume-engine/docs"
	"github.com/melegattip/financial-resume-engine/internal/config"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/http"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
	"github.com/melegattip/financial-resume-engine/internal/usecases/categories"
	"github.com/melegattip/financial-resume-engine/internal/usecases/reports"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/create"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/delete"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/get"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/list"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions/update"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Financial Resume Engine API
// @version         1.0
// @description     API para gestionar transacciones financieras y generar reportes
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @tag.name transactions
// @tag.description Operaciones relacionadas con transacciones financieras
// @tag.name categories
// @tag.description Operaciones relacionadas con categorías de transacciones
// @tag.name reports
// @tag.description Operaciones relacionadas con reportes financieros

// @Summary      Crear una nueva transacción
// @Description  Crea una nueva transacción financiera
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        transaction body create.CreateTransactionRequest true "Datos de la transacción"
// @Success      201 {object} create.CreateTransactionResponse
// @Failure      400 {object} http.ErrorResponse
// @Failure      401 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /transactions [post]
// @Security     BearerAuth
func handleCreateTransaction(c *gin.Context, handler *create.CreateTransaction) {
	if err := handler.Handle(c); err != nil {
		http.HandleError(c, err)
	}
}

// @Summary      Listar todas las transacciones
// @Description  Obtiene una lista de todas las transacciones
// @Tags         transactions
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Success      200 {array} transactions.TransactionModel
// @Failure      401 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /transactions [get]
// @Security     BearerAuth
func handleListTransactions(c *gin.Context, handler *list.ListTransactions) {
	if err := handler.Handle(c); err != nil {
		http.HandleError(c, err)
	}
}

// @Summary      Obtener una transacción específica
// @Description  Obtiene los detalles de una transacción por su ID
// @Tags         transactions
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        id path string true "ID de la transacción"
// @Success      200 {object} transactions.TransactionModel
// @Failure      401 {object} http.ErrorResponse
// @Failure      404 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /transactions/{id} [get]
// @Security     BearerAuth
func handleGetTransaction(c *gin.Context, handler *get.GetTransaction) {
	if err := handler.Handle(c); err != nil {
		http.HandleError(c, err)
	}
}

// @Summary      Actualizar una transacción
// @Description  Actualiza los datos de una transacción existente
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        id path string true "ID de la transacción"
// @Param        transaction body create.CreateTransactionRequest true "Datos actualizados de la transacción"
// @Success      200 {object} transactions.TransactionModel
// @Failure      400 {object} http.ErrorResponse
// @Failure      401 {object} http.ErrorResponse
// @Failure      404 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /transactions/{id} [patch]
// @Security     BearerAuth
func handleUpdateTransaction(c *gin.Context, handler *update.UpdateTransaction) {
	if err := handler.Handle(c); err != nil {
		http.HandleError(c, err)
	}
}

// @Summary      Eliminar una transacción
// @Description  Elimina una transacción por su ID
// @Tags         transactions
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        id path string true "ID de la transacción"
// @Success      204 "No Content"
// @Failure      401 {object} http.ErrorResponse
// @Failure      404 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /transactions/{id} [delete]
// @Security     BearerAuth
func handleDeleteTransaction(c *gin.Context, handler *delete.DeleteTransaction) {
	if err := handler.Handle(c); err != nil {
		http.HandleError(c, err)
	}
}

// @Summary      Crear una nueva categoría
// @Description  Crea una nueva categoría para transacciones
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        category body categories.CreateCategoryRequest true "Datos de la categoría"
// @Success      201 {object} categories.Category
// @Failure      400 {object} http.ErrorResponse
// @Failure      401 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /categories [post]
// @Security     BearerAuth
func handleCreateCategory(c *gin.Context, handler *categories.CategoryHandler) {
	handler.HandleCreateCategory(c)
}

// @Summary      Listar todas las categorías
// @Description  Obtiene una lista de todas las categorías
// @Tags         categories
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Success      200 {array} categories.Category
// @Failure      401 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /categories [get]
// @Security     BearerAuth
func handleListCategories(c *gin.Context, handler *categories.CategoryHandler) {
	handler.HandleGetCategories(c)
}

// @Summary      Actualizar una categoría
// @Description  Actualiza los datos de una categoría existente
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        id path string true "ID de la categoría"
// @Param        category body categories.UpdateCategoryRequest true "Datos actualizados de la categoría"
// @Success      200 {object} categories.Category
// @Failure      400 {object} http.ErrorResponse
// @Failure      401 {object} http.ErrorResponse
// @Failure      404 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /categories/{id} [patch]
// @Security     BearerAuth
func handleUpdateCategory(c *gin.Context, handler *categories.CategoryHandler) {
	handler.HandleUpdateCategory(c)
}

// @Summary      Eliminar una categoría
// @Description  Elimina una categoría por su ID
// @Tags         categories
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        id path string true "ID de la categoría"
// @Success      204 "No Content"
// @Failure      401 {object} http.ErrorResponse
// @Failure      404 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /categories/{id} [delete]
// @Security     BearerAuth
func handleDeleteCategory(c *gin.Context, handler *categories.CategoryHandler) {
	handler.HandleDeleteCategory(c)
}

// @Summary      Generar reporte financiero
// @Description  Genera un reporte financiero para un período específico
// @Tags         reports
// @Produce      json
// @Param        x-caller-id header string true "ID del llamador"
// @Param        start_date query string true "Fecha de inicio (YYYY-MM-DD)"
// @Param        end_date query string true "Fecha de fin (YYYY-MM-DD)"
// @Success      200 {object} reports.FinancialReport
// @Failure      400 {object} http.ErrorResponse
// @Failure      401 {object} http.ErrorResponse
// @Failure      500 {object} http.ErrorResponse
// @Router       /reports/financial [get]
// @Security     BearerAuth
func handleGenerateReport(c *gin.Context, handler *reports.ReportHandler) {
	handler.HandleGenerateReport(c)
}

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

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "x-caller-id"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 horas
	}))

	// Middleware para Swagger
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, x-caller-id")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "43200") // 12 horas

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Configurar Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))

	// Crear grupo de rutas API
	api := router.Group("/api/v1")
	{
		// Middleware para todas las rutas API
		api.Use(func(c *gin.Context) {
			c.Header("Content-Type", "application/json")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, x-caller-id")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Next()
		})

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
		api.POST("/transactions", func(c *gin.Context) {
			handleCreateTransaction(c, createTransactionHandler)
		})
		api.GET("/transactions", func(c *gin.Context) {
			handleListTransactions(c, listTransactionHandler)
		})
		api.GET("/transactions/:id", func(c *gin.Context) {
			handleGetTransaction(c, getTransactionHandler)
		})
		api.PATCH("/transactions/:id", func(c *gin.Context) {
			handleUpdateTransaction(c, updateTransactionHandler)
		})
		api.DELETE("/transactions/:id", func(c *gin.Context) {
			handleDeleteTransaction(c, deleteTransactionHandler)
		})

		// Rutas para categorías
		api.POST("/categories", func(c *gin.Context) {
			handleCreateCategory(c, categoryHandler)
		})
		api.GET("/categories", func(c *gin.Context) {
			handleListCategories(c, categoryHandler)
		})
		api.PATCH("/categories/:id", func(c *gin.Context) {
			handleUpdateCategory(c, categoryHandler)
		})
		api.DELETE("/categories/:id", func(c *gin.Context) {
			handleDeleteCategory(c, categoryHandler)
		})

		// Rutas para reportes
		api.GET("/reports/financial", func(c *gin.Context) {
			handleGenerateReport(c, reportHandler)
		})
	}

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
