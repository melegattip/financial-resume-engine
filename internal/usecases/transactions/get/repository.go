package get

import (
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Get(userID string, transactionID string) (*transactions.TransactionModel, error) {
	var transaction transactions.TransactionModel
	result := r.db.Where("user_id = ? AND id = ?", userID, transactionID).First(&transaction)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewResourceNotFound("Transaction not found")
		}
		return nil, result.Error
	}
	return &transaction, nil
}
