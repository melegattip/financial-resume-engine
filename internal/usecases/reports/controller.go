package reports

import (
	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{}
}

func (h *ReportHandler) HandleGenerateReport(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Not implemented yet",
	})
}
