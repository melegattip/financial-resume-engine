package list

import (
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) List(userID string) ([]transactions.TransactionModel, error) {
	var transactions []transactions.TransactionModel
	result := r.db.Where("user_id = ?", userID).Find(&transactions)
	return transactions, result.Error
}
