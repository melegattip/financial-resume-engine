package get

import (
	"context"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

type Get interface {
	Execute(ctx context.Context, userID string, transactionID string) (*GetTransactionResponse, error)
}

type GetImp struct {
	TransactionRepository *TransactionRepository
}

func (s *GetImp) Execute(ctx context.Context, userID string, transactionID string) (*GetTransactionResponse, error) {
	logger.Info(ctx, "Getting transaction", logs.Tags{"userID": userID, "transactionID": transactionID})

	transaction, err := s.TransactionRepository.Get(userID, transactionID)
	if err != nil {
		logger.Error(ctx, err, "Error getting transaction", logs.Tags{"error": err.Error(), "userID": userID, "transactionID": transactionID})
		return nil, err
	}

	responseBuilder := NewGetTransactionResponseBuilder()
	response := responseBuilder.
		SetID(transaction.ID).
		SetTypeID(transaction.TypeID).
		SetDescription(transaction.Description).
		SetAmount(transaction.Amount).
		SetPayed(transaction.Payed).
		SetCategory(transaction.CategoryID).
		SetCreatedAt(transaction.CreatedAt.Format(time.RFC3339)).
		SetUpdatedAt(transaction.UpdatedAt.Format(time.RFC3339))

	if !transaction.ExpiryDate.IsZero() {
		response.SetExpiryDate(transaction.ExpiryDate.Format(time.RFC3339))
	}

	return response.Build(), nil
}

func NewGetTransaction(repo *TransactionRepository) *GetImp {
	return &GetImp{TransactionRepository: repo}
}
