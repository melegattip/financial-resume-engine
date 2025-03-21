package create

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

func (r *TransactionRepository) Create(transaction *transactions.TransactionModel) error {
	result := r.db.Create(transaction)
	return result.Error
}

func (r *TransactionRepository) List(userID string) ([]transactions.TransactionModel, error) {
	var transactions []transactions.TransactionModel
	result := r.db.Where("user_id = ?", userID).Find(&transactions)
	return transactions, result.Error
}

func (r *TransactionRepository) Get(userID string, id string, transaction *transactions.TransactionModel) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, id).First(transaction)
	return result.Error
}

func (r *TransactionRepository) Update(userID string, id string, updates map[string]interface{}) error {
	result := r.db.Model(&transactions.TransactionModel{}).Where("user_id = ? AND id = ?", userID, id).Updates(updates)
	return result.Error
}

func (r *TransactionRepository) Delete(userID string, id string) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&transactions.TransactionModel{})
	return result.Error
}
