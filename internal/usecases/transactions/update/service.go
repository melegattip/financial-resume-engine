package update

import (
	"context"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	"github.com/melegattip/financial-resume-engine/internal/infrastructure/logger"
)

type Update interface {
	Execute(ctx context.Context, userID string, transactionID string, request *UpdateTransactionRequest) (*UpdateTransactionResponse, error)
}

type UpdateImp struct {
	TransactionRepository *TransactionRepository
}

func (s *UpdateImp) Execute(ctx context.Context, userID string, transactionID string, request *UpdateTransactionRequest) (*UpdateTransactionResponse, error) {
	logger.Info(ctx, "Updating transaction", logs.Tags{"userID": userID, "transactionID": transactionID})

	// Obtener la transacción existente
	transaction, err := s.TransactionRepository.Get(userID, transactionID)
	if err != nil {
		return nil, err
	}

	// Actualizar los campos que vienen en el request
	if request.TypeID != "" {
		transaction.TypeID = request.TypeID
	}
	if request.Description != "" {
		transaction.Description = request.Description
	}
	if request.Amount != 0 {
		transaction.Amount = request.Amount
	}
	if request.Category != "" {
		transaction.CategoryID = request.Category
	}
	transaction.Payed = request.Payed
	transaction.UpdatedAt = time.Now().UTC()

	if request.ExpiryDate != "" {
		expiryDate, err := time.Parse(time.RFC3339, request.ExpiryDate)
		if err != nil {
			return nil, err
		}
		transaction.ExpiryDate = expiryDate
	}

	// Guardar los cambios
	if err := s.TransactionRepository.Update(transaction); err != nil {
		return nil, err
	}

	// Construir la respuesta
	responseBuilder := NewUpdateTransactionResponseBuilder()
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

func NewUpdateTransaction(repo *TransactionRepository) *UpdateImp {
	return &UpdateImp{TransactionRepository: repo}
}
