package expenses

import (
	"context"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	baseRepo "github.com/melegattip/financial-resume-engine/internal/core/repository"
)

// ExpenseService define las operaciones disponibles para el servicio de gastos
type ExpenseService interface {
	CreateExpense(ctx context.Context, request *CreateExpenseRequest) (*CreateExpenseResponse, error)
	GetExpense(ctx context.Context, userID string, expenseID string) (*GetExpenseResponse, error)
	ListExpenses(ctx context.Context, userID string) (*ListExpensesResponse, error)
	ListUnpaidExpenses(ctx context.Context, userID string) (*ListExpensesResponse, error)
	ListExpensesByDueDate(ctx context.Context, userID string) (*ListExpensesResponse, error)
	UpdateExpense(ctx context.Context, userID string, expenseID string, request *UpdateExpenseRequest) (*UpdateExpenseResponse, error)
	MarkAsPaid(ctx context.Context, userID string, expenseID string) (*MarkAsPaidResponse, error)
	DeleteExpense(ctx context.Context, userID string, expenseID string) error
}

// ExpenseServiceImpl implementa ExpenseService
type ExpenseServiceImpl struct {
	repository baseRepo.ExpenseRepository
}

func NewExpenseService(repository baseRepo.ExpenseRepository) ExpenseService {
	return &ExpenseServiceImpl{
		repository: repository,
	}
}

func (s *ExpenseServiceImpl) CreateExpense(ctx context.Context, request *CreateExpenseRequest) (*CreateExpenseResponse, error) {
	dueDate, err := time.Parse("2006-01-02T15:04:05Z07:00", request.DueDate)
	if err != nil {
		return nil, errors.NewBadRequest("Invalid due date format")
	}

	expense := domain.NewExpenseBuilder().
		SetID(domain.NewID()).
		SetUserID(request.UserID).
		SetAmount(request.Amount).
		SetDescription(request.Description).
		SetCategory(request.Category).
		SetPaid(request.Paid).
		SetDueDate(dueDate).
		Build()

	if err := s.repository.Create(expense); err != nil {
		return nil, err
	}

	return &CreateExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		Amount:      expense.Amount,
		Description: expense.Description,
		Category:    expense.Category,
		Paid:        expense.Paid,
		DueDate:     expense.DueDate.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   expense.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *ExpenseServiceImpl) GetExpense(ctx context.Context, userID string, expenseID string) (*GetExpenseResponse, error) {
	expense, err := s.repository.Get(userID, expenseID)
	if err != nil {
		return nil, err
	}

	return &GetExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		Amount:      expense.Amount,
		Description: expense.Description,
		Category:    expense.Category,
		Paid:        expense.Paid,
		DueDate:     expense.DueDate.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   expense.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *ExpenseServiceImpl) ListExpenses(ctx context.Context, userID string) (*ListExpensesResponse, error) {
	expenses, err := s.repository.List(userID)
	if err != nil {
		return nil, err
	}

	var response []GetExpenseResponse
	for _, expense := range expenses {
		response = append(response, GetExpenseResponse{
			ID:          expense.ID,
			UserID:      expense.UserID,
			Amount:      expense.Amount,
			Description: expense.Description,
			Category:    expense.Category,
			Paid:        expense.Paid,
			DueDate:     expense.DueDate.Format("2006-01-02T15:04:05Z07:00"),
			CreatedAt:   expense.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &ListExpensesResponse{
		Expenses: response,
	}, nil
}

func (s *ExpenseServiceImpl) ListUnpaidExpenses(ctx context.Context, userID string) (*ListExpensesResponse, error) {
	expenses, err := s.repository.ListUnpaid(userID)
	if err != nil {
		return nil, err
	}

	var response []GetExpenseResponse
	for _, expense := range expenses {
		response = append(response, GetExpenseResponse{
			ID:          expense.ID,
			UserID:      expense.UserID,
			Amount:      expense.Amount,
			Description: expense.Description,
			Category:    expense.Category,
			Paid:        expense.Paid,
			DueDate:     expense.DueDate.Format("2006-01-02T15:04:05Z07:00"),
			CreatedAt:   expense.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &ListExpensesResponse{
		Expenses: response,
	}, nil
}

func (s *ExpenseServiceImpl) ListExpensesByDueDate(ctx context.Context, userID string) (*ListExpensesResponse, error) {
	expenses, err := s.repository.ListByDueDate(userID)
	if err != nil {
		return nil, err
	}

	var response []GetExpenseResponse
	for _, expense := range expenses {
		response = append(response, GetExpenseResponse{
			ID:          expense.ID,
			UserID:      expense.UserID,
			Amount:      expense.Amount,
			Description: expense.Description,
			Category:    expense.Category,
			Paid:        expense.Paid,
			DueDate:     expense.DueDate.Format("2006-01-02T15:04:05Z07:00"),
			CreatedAt:   expense.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &ListExpensesResponse{
		Expenses: response,
	}, nil
}

func (s *ExpenseServiceImpl) UpdateExpense(ctx context.Context, userID string, expenseID string, request *UpdateExpenseRequest) (*UpdateExpenseResponse, error) {
	expense, err := s.repository.Get(userID, expenseID)
	if err != nil {
		return nil, err
	}

	if request.Amount != 0 {
		expense.Amount = request.Amount
	}
	if request.Description != "" {
		expense.Description = request.Description
	}
	if request.Category != "" {
		expense.Category = request.Category
	}
	expense.Paid = request.Paid
	if request.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02T15:04:05Z07:00", request.DueDate)
		if err != nil {
			return nil, errors.NewBadRequest("Invalid due date format")
		}
		expense.DueDate = dueDate
	}
	expense.UpdatedAt = time.Now()

	if err := s.repository.Update(expense); err != nil {
		return nil, err
	}

	return &UpdateExpenseResponse{
		ID:          expense.ID,
		UserID:      expense.UserID,
		Amount:      expense.Amount,
		Description: expense.Description,
		Category:    expense.Category,
		Paid:        expense.Paid,
		DueDate:     expense.DueDate.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   expense.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *ExpenseServiceImpl) MarkAsPaid(ctx context.Context, userID string, expenseID string) (*MarkAsPaidResponse, error) {
	if err := s.repository.MarkAsPaid(userID, expenseID); err != nil {
		return nil, err
	}

	expense, err := s.repository.Get(userID, expenseID)
	if err != nil {
		return nil, err
	}

	return &MarkAsPaidResponse{
		ID:        expense.ID,
		Paid:      expense.Paid,
		UpdatedAt: expense.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *ExpenseServiceImpl) DeleteExpense(ctx context.Context, userID string, expenseID string) error {
	_, err := s.repository.Get(userID, expenseID)
	if err != nil {
		return err
	}

	return s.repository.Delete(userID, expenseID)
}
