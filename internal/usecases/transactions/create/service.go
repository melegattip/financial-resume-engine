package create

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions"
)

type Create interface {
	Execute(ctx context.Context, transaction *CreateTransactionRequest) (*CreateTransactionResponse, error)
}

type CreateImp struct {
	TransactionRepository *TransactionRepository
}

func (s *CreateImp) Execute(ctx context.Context, transaction *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	logger.Info(context.Background(), "Creating transaction", logs.Tags{"userID": transaction.UserID, "transaction": transaction})

	transactionCreatedAt := time.Now().UTC()
	transactionUpdatedAt := transactionCreatedAt
	transactionID := "tx_" + uuid.New().String()[:8]

	transactionModel := transactions.NewTransactionBuilder().
		SetID(transactionID).
		SetUserID(transaction.UserID).
		SetTypeID(transaction.TypeID).
		SetCategoryID(transaction.CategoryID).
		SetAmount(transaction.Amount).
		SetPayed(transaction.Payed).
		SetExpiryDate(stringToTime(transaction.ExpiryDate)).
		SetCreatedAt(transactionCreatedAt).
		SetUpdatedAt(transactionUpdatedAt).
		Build()

	err := s.TransactionRepository.Create(transactionModel)
	if err != nil {
		return nil, err
	}

	response := &CreateTransactionResponse{
		TransactionID: transactionID,
		CreatedAt:     transactionModel.CreatedAt.Format(time.RFC3339),
	}

	return response, nil
}

type ListTransactions struct {
	TransactionRepository *TransactionRepository
}

func NewListTransactions(repo *TransactionRepository) *ListTransactions {
	return &ListTransactions{TransactionRepository: repo}
}

func (s *ListTransactions) Execute(userID string) ([]transactions.TransactionModel, error) {
	return s.TransactionRepository.List(userID)
}

type GetTransaction struct {
	TransactionRepository *TransactionRepository
}

func NewGetTransaction(repo *TransactionRepository) *GetTransaction {
	return &GetTransaction{TransactionRepository: repo}
}

func (s *GetTransaction) Execute(userID string, id string) (*transactions.TransactionModel, error) {
	var transaction transactions.TransactionModel
	if err := s.TransactionRepository.Get(userID, id, &transaction); err != nil {
		return nil, err
	}
	return &transaction, nil
}

type UpdateTransaction struct {
	TransactionRepository *TransactionRepository
}

func NewUpdateTransaction(repo *TransactionRepository) *UpdateTransaction {
	return &UpdateTransaction{TransactionRepository: repo}
}

func (s *UpdateTransaction) Execute(userID string, id string, updates map[string]interface{}) error {
	return s.TransactionRepository.Update(userID, id, updates)
}

type DeleteTransaction struct {
	TransactionRepository *TransactionRepository
}

func NewDeleteTransaction(repo *TransactionRepository) *DeleteTransaction {
	return &DeleteTransaction{TransactionRepository: repo}
}

func (s *DeleteTransaction) Execute(userID string, id string) error {
	return s.TransactionRepository.Delete(userID, id)
}

func stringToTime(date string) time.Time {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}
	}
	return t
}

func NewCreateTransaction(repo *TransactionRepository) *CreateImp {
	return &CreateImp{TransactionRepository: repo}
}
