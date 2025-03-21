package update

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

func (r *TransactionRepository) Get(userID string, transactionID string) (*transactions.TransactionModel, error) {
	var transaction transactions.TransactionModel
	if err := r.db.Where("user_id = ? AND id = ?", userID, transactionID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) Update(transaction *transactions.TransactionModel) error {
	return r.db.Save(transaction).Error
}
