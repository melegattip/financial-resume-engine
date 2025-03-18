package reports

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FinancialReport struct {
	TotalIncome   float64 `json:"total_income"`
	TotalExpenses float64 `json:"total_expenses"`
	NetBalance    float64 `json:"net_balance"`
	PeriodStart   string  `json:"period_start"`
	PeriodEnd     string  `json:"period_end"`
}

func GenerateFinancialReport(c *gin.Context) {
	// TODO: Implementar lógica de generación de reporte
	report := FinancialReport{
		TotalIncome:   0,
		TotalExpenses: 0,
		NetBalance:    0,
	}
	c.JSON(http.StatusOK, report)
}
