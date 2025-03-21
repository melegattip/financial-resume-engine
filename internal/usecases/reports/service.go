package reports

import (
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

type GenerateFinancialReport struct {
	repository ReportRepository
}

func NewGenerateFinancialReport(repository ReportRepository) *GenerateFinancialReport {
	return &GenerateFinancialReport{
		repository: repository,
	}
}

func (s *GenerateFinancialReport) Execute(startDate, endDate time.Time, userID string) (*FinancialReport, error) {
	transactions, err := s.repository.GetTransactions(startDate, endDate, userID)
	if err != nil {
		logger.Error(nil, err, logs.ErrorCreatingCategory.Message, logs.Tags{
			"start_date": startDate,
			"end_date":   endDate,
			"user_id":    userID,
		})
		return nil, err
	}

	builder := NewFinancialReportBuilder()
	report := builder.
		WithStartDate(startDate).
		WithEndDate(endDate).
		WithTransactions(transactions).
		Build()

	return report, nil
}
