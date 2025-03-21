package reports

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      string    `gorm:"not null"`
	Amount      float64   `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
}

// FinancialReport representa el reporte financiero
type FinancialReport struct {
	StartDate     time.Time     `json:"start_date"`
	EndDate       time.Time     `json:"end_date"`
	Transactions  []Transaction `json:"transactions"`
	TotalIncome   float64       `json:"total_income"`
	TotalExpenses float64       `json:"total_expenses"`
	NetBalance    float64       `json:"net_balance"`
}

// GenerateReportRequest representa los parámetros para generar un reporte
type GenerateReportRequest struct {
	StartDate string `form:"start_date" binding:"required"`
	EndDate   string `form:"end_date" binding:"required"`
}

// Builder pattern
type FinancialReportBuilder struct {
	report *FinancialReport
}

func NewFinancialReportBuilder() *FinancialReportBuilder {
	return &FinancialReportBuilder{
		report: &FinancialReport{},
	}
}

func (b *FinancialReportBuilder) WithStartDate(date time.Time) *FinancialReportBuilder {
	b.report.StartDate = date
	return b
}

func (b *FinancialReportBuilder) WithEndDate(date time.Time) *FinancialReportBuilder {
	b.report.EndDate = date
	return b
}

func (b *FinancialReportBuilder) WithTransactions(transactions []Transaction) *FinancialReportBuilder {
	b.report.Transactions = transactions
	var totalIncome, totalExpenses float64
	for _, t := range transactions {
		if t.Amount > 0 {
			totalIncome += t.Amount
		} else {
			totalExpenses += -t.Amount
		}
	}
	b.report.TotalIncome = totalIncome
	b.report.TotalExpenses = totalExpenses
	b.report.NetBalance = totalIncome - totalExpenses
	return b
}

func (b *FinancialReportBuilder) Build() *FinancialReport {
	return b.report
}
