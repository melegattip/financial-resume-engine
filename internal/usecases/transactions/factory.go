package transactions

import (
	"context"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/usecases/expenses"
	"github.com/melegattip/financial-resume-engine/internal/usecases/incomes"
)

// TransactionType define el tipo de transacción
type TransactionType string

const (
	IncomeType  TransactionType = "income"
	ExpenseType TransactionType = "expense"
)

// TransactionFactory es la interfaz que define el factory para crear transacciones
type TransactionFactory interface {
	CreateTransaction(ctx context.Context, userID string, amount float64, description, category string, dueDate *time.Time) (Transaction, error)
	GetTransaction(ctx context.Context, userID string, transactionID string, transactionType TransactionType) (Transaction, error)
	ListTransactions(ctx context.Context, userID string, transactionType TransactionType) ([]Transaction, error)
	UpdateTransaction(ctx context.Context, userID string, transactionID string, amount float64, description, category string, dueDate *time.Time, transactionType TransactionType) (Transaction, error)
	DeleteTransaction(ctx context.Context, userID string, transactionID string, transactionType TransactionType) error
}

// Transaction es la interfaz que deben implementar todos los tipos de transacciones
type Transaction interface {
	GetID() string
	GetUserID() string
	GetAmount() float64
	GetDescription() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// TransactionFactoryImpl implementa TransactionFactory
type TransactionFactoryImpl struct {
	incomeService  incomes.IncomeService
	expenseService expenses.ExpenseService
}

// NewTransactionFactory crea una nueva instancia de TransactionFactory
func NewTransactionFactory(incomeService incomes.IncomeService, expenseService expenses.ExpenseService) TransactionFactory {
	return &TransactionFactoryImpl{
		incomeService:  incomeService,
		expenseService: expenseService,
	}
}

func (f *TransactionFactoryImpl) CreateTransaction(ctx context.Context, userID string, amount float64, description, category string, dueDate *time.Time) (Transaction, error) {
	if dueDate != nil {
		// Si tiene fecha de vencimiento, es un gasto
		request := &expenses.CreateExpenseRequest{
			UserID:      userID,
			Amount:      amount,
			Description: description,
			Category:    category,
			DueDate:     dueDate.Format("2006-01-02"),
		}
		response, err := f.expenseService.CreateExpense(ctx, request)
		if err != nil {
			return nil, err
		}
		return toTransaction(response), nil
	}

	// Si no tiene fecha de vencimiento, es un ingreso
	request := &incomes.CreateIncomeRequest{
		UserID:      userID,
		Amount:      amount,
		Description: description,
		Category:    category,
	}
	response, err := f.incomeService.CreateIncome(ctx, request)
	if err != nil {
		return nil, err
	}
	return toTransaction(response), nil
}

func (f *TransactionFactoryImpl) GetTransaction(ctx context.Context, userID string, transactionID string, transactionType TransactionType) (Transaction, error) {
	switch transactionType {
	case IncomeType:
		response, err := f.incomeService.GetIncome(ctx, userID, transactionID)
		if err != nil {
			return nil, err
		}
		return toTransaction(response), nil
	case ExpenseType:
		response, err := f.expenseService.GetExpense(ctx, userID, transactionID)
		if err != nil {
			return nil, err
		}
		return toTransaction(response), nil
	default:
		return nil, domain.ErrInvalidTransactionType
	}
}

func (f *TransactionFactoryImpl) ListTransactions(ctx context.Context, userID string, transactionType TransactionType) ([]Transaction, error) {
	switch transactionType {
	case IncomeType:
		response, err := f.incomeService.ListIncomes(ctx, userID)
		if err != nil {
			return nil, err
		}
		return toTransactionSlice(response.Incomes), nil
	case ExpenseType:
		response, err := f.expenseService.ListExpenses(ctx, userID)
		if err != nil {
			return nil, err
		}
		return toTransactionSlice(response.Expenses), nil
	default:
		return nil, domain.ErrInvalidTransactionType
	}
}

func (f *TransactionFactoryImpl) UpdateTransaction(ctx context.Context, userID string, transactionID string, amount float64, description, category string, dueDate *time.Time, transactionType TransactionType) (Transaction, error) {
	switch transactionType {
	case IncomeType:
		request := &incomes.UpdateIncomeRequest{
			Amount:      amount,
			Description: description,
			Category:    category,
		}
		response, err := f.incomeService.UpdateIncome(ctx, userID, transactionID, request)
		if err != nil {
			return nil, err
		}
		return toTransaction(response), nil
	case ExpenseType:
		request := &expenses.UpdateExpenseRequest{
			Amount:      amount,
			Description: description,
			Category:    category,
		}
		if dueDate != nil {
			request.DueDate = dueDate.Format("2006-01-02")
		}
		response, err := f.expenseService.UpdateExpense(ctx, userID, transactionID, request)
		if err != nil {
			return nil, err
		}
		return toTransaction(response), nil
	default:
		return nil, domain.ErrInvalidTransactionType
	}
}

func (f *TransactionFactoryImpl) DeleteTransaction(ctx context.Context, userID string, transactionID string, transactionType TransactionType) error {
	switch transactionType {
	case IncomeType:
		return f.incomeService.DeleteIncome(ctx, userID, transactionID)
	case ExpenseType:
		return f.expenseService.DeleteExpense(ctx, userID, transactionID)
	default:
		return domain.ErrInvalidTransactionType
	}
}

// Helper functions para convertir las respuestas a la interfaz Transaction
func toTransaction(response interface{}) Transaction {
	switch v := response.(type) {
	case *incomes.CreateIncomeResponse:
		return &TransactionModel{
			ID:          v.ID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   parseTime(v.CreatedAt),
			UpdatedAt:   parseTime(v.UpdatedAt),
		}
	case *incomes.GetIncomeResponse:
		return &TransactionModel{
			ID:          v.ID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   parseTime(v.CreatedAt),
			UpdatedAt:   parseTime(v.UpdatedAt),
		}
	case *incomes.UpdateIncomeResponse:
		return &TransactionModel{
			ID:          v.ID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   parseTime(v.CreatedAt),
			UpdatedAt:   parseTime(v.UpdatedAt),
		}
	case *expenses.CreateExpenseResponse:
		return &TransactionModel{
			ID:          v.ID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   parseTime(v.CreatedAt),
			UpdatedAt:   parseTime(v.UpdatedAt),
		}
	case *expenses.GetExpenseResponse:
		return &TransactionModel{
			ID:          v.ID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   parseTime(v.CreatedAt),
			UpdatedAt:   parseTime(v.UpdatedAt),
		}
	case *expenses.UpdateExpenseResponse:
		return &TransactionModel{
			ID:          v.ID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Description: v.Description,
			Category:    v.Category,
			CreatedAt:   parseTime(v.CreatedAt),
			UpdatedAt:   parseTime(v.UpdatedAt),
		}
	default:
		return nil
	}
}

func toTransactionSlice(responses interface{}) []Transaction {
	var transactions []Transaction

	switch v := responses.(type) {
	case []incomes.GetIncomeResponse:
		for _, income := range v {
			transactions = append(transactions, toTransaction(&income))
		}
	case []expenses.GetExpenseResponse:
		for _, expense := range v {
			transactions = append(transactions, toTransaction(&expense))
		}
	}

	return transactions
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z07:00", timeStr)
	return t
}
