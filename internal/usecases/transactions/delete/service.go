package delete

import (
	"context"

	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

type Delete interface {
	Execute(ctx context.Context, userID string, transactionID string) (*DeleteTransactionResponse, error)
}

type DeleteImp struct {
	TransactionRepository *TransactionRepository
}

func (s *DeleteImp) Execute(ctx context.Context, userID string, transactionID string) (*DeleteTransactionResponse, error) {
	logger.Info(ctx, "Deleting transaction", logs.Tags{"userID": userID, "transactionID": transactionID})

	// Obtener la transacción existente
	transaction, err := s.TransactionRepository.Get(userID, transactionID)
	if err != nil {
		return nil, err
	}

	// Eliminar la transacción
	if err := s.TransactionRepository.Delete(transaction); err != nil {
		return nil, err
	}

	return &DeleteTransactionResponse{
		Message: "Transaction deleted successfully",
	}, nil
}

func NewDeleteTransaction(repo *TransactionRepository) *DeleteImp {
	return &DeleteImp{TransactionRepository: repo}
}
