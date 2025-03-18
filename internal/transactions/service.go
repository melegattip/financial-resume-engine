package transactions

import (
	"time"

	"github.com/google/uuid"
	"github.com/melegattip/financial-resume-engine/internal/models"
)

type CreateTransaction struct {
	TransactionRepository *TransactionRepository
}

func NewCreateTransaction(repo *TransactionRepository) *CreateTransaction {
	return &CreateTransaction{TransactionRepository: repo}
}

func (s *CreateTransaction) Execute(userID string, transaction *models.TransactionModel) (*models.TransactionResponse, error) {
	// Generar ID único para la transacción
	transactionID := "tx_" + uuid.New().String()[:8]

	// Establecer fecha de creación y user_id
	transaction.CreatedAt = time.Now().UTC()
	transaction.UpdatedAt = transaction.CreatedAt
	transaction.UserID = userID

	// Crear la transacción en el repositorio
	err := s.TransactionRepository.Create(transactionID, transaction)
	if err != nil {
		return nil, err
	}

	// Construir respuesta
	response := models.NewTransactionResponseBuilder().
		SetTransactionID(transactionID).
		SetCreatedAt(transaction.CreatedAt.Format(time.RFC3339)).
		Build()

	return response, nil
}
