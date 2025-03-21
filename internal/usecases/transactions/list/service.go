package list

import (
	"context"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

type List interface {
	Execute(ctx context.Context, userID string) (*ListTransactionsResponse, error)
}

type ListImp struct {
	TransactionRepository *TransactionRepository
}

func (s *ListImp) Execute(ctx context.Context, userID string) (*ListTransactionsResponse, error) {
	logger.Info(ctx, "Listing transactions", logs.Tags{"userID": userID})

	transactions, err := s.TransactionRepository.List(userID)
	if err != nil {
		return nil, err
	}

	responseBuilder := NewListTransactionsResponseBuilder()
	var transactionResponses []TransactionResponse

	for _, t := range transactions {
		transactionResponse := TransactionResponse{
			ID:          t.ID,
			TypeID:      t.TypeID,
			Description: t.Description,
			Amount:      t.Amount,
			Payed:       t.Payed,
			Category:    t.CategoryID,
			CreatedAt:   t.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   t.UpdatedAt.Format(time.RFC3339),
		}

		if !t.ExpiryDate.IsZero() {
			transactionResponse.ExpiryDate = t.ExpiryDate.Format(time.RFC3339)
		}

		transactionResponses = append(transactionResponses, transactionResponse)
	}

	response := responseBuilder.SetTransactions(transactionResponses).Build()
	return response, nil
}

func NewListTransactions(repo *TransactionRepository) *ListImp {
	return &ListImp{TransactionRepository: repo}
}
