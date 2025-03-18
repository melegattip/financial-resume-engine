package transactions

import (
	"github.com/melegattip/financial-resume-engine/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transactionID string, transaction *models.TransactionModel) error {
	transaction.ID = transactionID
	result := r.db.Create(transaction)
	return result.Error
}

func (r *TransactionRepository) List(userID string) ([]models.TransactionModel, error) {
	var transactions []models.TransactionModel
	result := r.db.Where("user_id = ?", userID).Find(&transactions)
	return transactions, result.Error
}

func (r *TransactionRepository) Update(userID string, transaction *models.TransactionModel) error {
	result := r.db.Where("user_id = ?", userID).Save(transaction)
	return result.Error
}

func (r *TransactionRepository) Delete(userID string, id string) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&models.TransactionModel{})
	return result.Error
}
