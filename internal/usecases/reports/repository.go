package reports

import (
	"time"

	"gorm.io/gorm"
)

// ReportRepository define la interfaz para el repositorio de reportes
type ReportRepository interface {
	GetTransactions(startDate, endDate time.Time, userID string) ([]Transaction, error)
}

type ReportRepositoryImpl struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &ReportRepositoryImpl{db: db}
}

func (r *ReportRepositoryImpl) GetTransactions(startDate, endDate time.Time, userID string) ([]Transaction, error) {
	var transactions []Transaction
	result := r.db.Where("created_at BETWEEN ? AND ? AND user_id = ?", startDate, endDate, userID).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
