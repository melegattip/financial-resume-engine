package docs

import (
	"github.com/gin-gonic/gin"
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

// Transaction representa una transacción financiera
type Transaction struct {
	UserID      string  `json:"user_id" example:"123"`
	TypeID      string  `json:"type_id" example:"income"`
	Description string  `json:"description" example:"Salario mensual"`
	Amount      float64 `json:"amount" example:"5000.00"`
	Payed       bool    `json:"payed" example:"true"`
	ExpiryDate  string  `json:"expiry_date" example:"2024-03-21"`
	CategoryID  string  `json:"category" example:"salary"`
}

// ErrorResponse representa una respuesta de error
type ErrorResponse struct {
	Error string `json:"error" example:"Bad Request"`
}

// @Summary      Crear una nueva transacción
// @Description  Crea una nueva transacción financiera
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        transaction body create.CreateTransactionRequest true "Datos de la transacción"
// @Success      201 {object} create.CreateTransactionResponse
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /transactions [post]
// @Security     BearerAuth
func CreateTransaction(c *gin.Context) {}

// @Summary      Listar todas las transacciones
// @Description  Obtiene una lista de todas las transacciones
// @Tags         transactions
// @Produce      json
// @Success      200 {array} model.TransactionModel
// @Failure      401 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /transactions [get]
// @Security     BearerAuth
func ListTransactions(c *gin.Context) {}

// @Summary      Obtener una transacción específica
// @Description  Obtiene los detalles de una transacción por su ID
// @Tags         transactions
// @Produce      json
// @Param        id path string true "ID de la transacción"
// @Success      200 {object} model.TransactionModel
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /transactions/{id} [get]
// @Security     BearerAuth
func GetTransaction(c *gin.Context) {}

// @Summary      Actualizar una transacción
// @Description  Actualiza los datos de una transacción existente
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id path string true "ID de la transacción"
// @Param        transaction body create.CreateTransactionRequest true "Datos actualizados de la transacción"
// @Success      200 {object} model.TransactionModel
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /transactions/{id} [patch]
// @Security     BearerAuth
func UpdateTransaction(c *gin.Context) {}

// @Summary      Eliminar una transacción
// @Description  Elimina una transacción por su ID
// @Tags         transactions
// @Produce      json
// @Param        id path string true "ID de la transacción"
// @Success      204 "No Content"
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /transactions/{id} [delete]
// @Security     BearerAuth
func DeleteTransaction(c *gin.Context) {}

// @Summary      Crear una nueva categoría
// @Description  Crea una nueva categoría para transacciones
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category body Category true "Datos de la categoría"
// @Success      201 {object} Category
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /categories [post]
// @Security     BearerAuth
func CreateCategory(c *gin.Context) {}

// @Summary      Listar todas las categorías
// @Description  Obtiene una lista de todas las categorías
// @Tags         categories
// @Produce      json
// @Success      200 {array} Category
// @Failure      401 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /categories [get]
// @Security     BearerAuth
func ListCategories(c *gin.Context) {}

// @Summary      Actualizar una categoría
// @Description  Actualiza los datos de una categoría existente
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "ID de la categoría"
// @Param        category body Category true "Datos actualizados de la categoría"
// @Success      200 {object} Category
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /categories/{id} [patch]
// @Security     BearerAuth
func UpdateCategory(c *gin.Context) {}

// @Summary      Eliminar una categoría
// @Description  Elimina una categoría por su ID
// @Tags         categories
// @Produce      json
// @Param        id path string true "ID de la categoría"
// @Success      204 "No Content"
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /categories/{id} [delete]
// @Security     BearerAuth
func DeleteCategory(c *gin.Context) {}

// @Summary      Generar reporte financiero
// @Description  Genera un reporte financiero para un período específico
// @Tags         reports
// @Produce      json
// @Param        start_date query string true "Fecha de inicio (YYYY-MM-DD)"
// @Param        end_date query string true "Fecha de fin (YYYY-MM-DD)"
// @Success      200 {object} Report
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /reports [get]
// @Security     BearerAuth
func GenerateReport(c *gin.Context) {}
