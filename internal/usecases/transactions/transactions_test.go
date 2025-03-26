package transactions

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/usecases/expenses"
	"github.com/melegattip/financial-resume-engine/internal/usecases/incomes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockIncomeService es un mock del servicio de ingresos
type MockIncomeService struct {
	mock.Mock
}

func (m *MockIncomeService) CreateIncome(ctx context.Context, request *incomes.CreateIncomeRequest) (*incomes.CreateIncomeResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*incomes.CreateIncomeResponse), args.Error(1)
}

func (m *MockIncomeService) GetIncome(ctx context.Context, userID string, incomeID string) (*incomes.GetIncomeResponse, error) {
	args := m.Called(ctx, userID, incomeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*incomes.GetIncomeResponse), args.Error(1)
}

func (m *MockIncomeService) ListIncomes(ctx context.Context, userID string) (*incomes.ListIncomesResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*incomes.ListIncomesResponse), args.Error(1)
}

func (m *MockIncomeService) UpdateIncome(ctx context.Context, userID string, incomeID string, request *incomes.UpdateIncomeRequest) (*incomes.UpdateIncomeResponse, error) {
	args := m.Called(ctx, userID, incomeID, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*incomes.UpdateIncomeResponse), args.Error(1)
}

func (m *MockIncomeService) DeleteIncome(ctx context.Context, userID string, incomeID string) error {
	args := m.Called(ctx, userID, incomeID)
	return args.Error(0)
}

// MockExpenseService es un mock del servicio de gastos
type MockExpenseService struct {
	mock.Mock
}

func (m *MockExpenseService) CreateExpense(ctx context.Context, request *expenses.CreateExpenseRequest) (*expenses.CreateExpenseResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.CreateExpenseResponse), args.Error(1)
}

func (m *MockExpenseService) GetExpense(ctx context.Context, userID string, expenseID string) (*expenses.GetExpenseResponse, error) {
	args := m.Called(ctx, userID, expenseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.GetExpenseResponse), args.Error(1)
}

func (m *MockExpenseService) ListExpenses(ctx context.Context, userID string) (*expenses.ListExpensesResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.ListExpensesResponse), args.Error(1)
}

func (m *MockExpenseService) ListUnpaidExpenses(ctx context.Context, userID string) (*expenses.ListExpensesResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.ListExpensesResponse), args.Error(1)
}

func (m *MockExpenseService) ListExpensesByDueDate(ctx context.Context, userID string) (*expenses.ListExpensesResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.ListExpensesResponse), args.Error(1)
}

func (m *MockExpenseService) UpdateExpense(ctx context.Context, userID string, expenseID string, request *expenses.UpdateExpenseRequest) (*expenses.UpdateExpenseResponse, error) {
	args := m.Called(ctx, userID, expenseID, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.UpdateExpenseResponse), args.Error(1)
}

func (m *MockExpenseService) MarkAsPaid(ctx context.Context, userID string, expenseID string) (*expenses.MarkAsPaidResponse, error) {
	args := m.Called(ctx, userID, expenseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*expenses.MarkAsPaidResponse), args.Error(1)
}

func (m *MockExpenseService) DeleteExpense(ctx context.Context, userID string, expenseID string) error {
	args := m.Called(ctx, userID, expenseID)
	return args.Error(0)
}

func TestCreateTransaction(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test crear un gasto
	dueDate := time.Now().AddDate(0, 1, 0)
	expenseRequest := &expenses.CreateExpenseRequest{
		UserID:      "user1",
		Amount:      100.0,
		Description: "Test expense",
		Category:    "food",
		Paid:        false,
		DueDate:     dueDate.Format("2006-01-02T15:04:05Z07:00"),
	}

	expenseResponse := &expenses.CreateExpenseResponse{
		ID:          "1",
		UserID:      "user1",
		Amount:      100.0,
		Description: "Test expense",
		Category:    "food",
		Paid:        false,
		DueDate:     dueDate.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	mockExpenseService.On("CreateExpense", context.Background(), expenseRequest).Return(expenseResponse, nil)

	result, err := factory.CreateTransaction(context.Background(), "user1", 100.0, "Test expense", "food", &dueDate)
	assert.NoError(t, err)
	assert.Equal(t, expenseResponse.ID, result.GetID())
	assert.Equal(t, expenseResponse.UserID, result.GetUserID())
	assert.Equal(t, expenseResponse.Amount, result.GetAmount())
	assert.Equal(t, expenseResponse.Description, result.GetDescription())
	assert.Equal(t, expenseResponse.Category, result.GetCategory())
	mockExpenseService.AssertExpectations(t)

	// Test crear un ingreso
	incomeRequest := &incomes.CreateIncomeRequest{
		UserID:      "user1",
		Amount:      200.0,
		Description: "Test income",
		Category:    "salary",
		Source:      "employer",
	}

	incomeResponse := &incomes.CreateIncomeResponse{
		ID:          "2",
		UserID:      "user1",
		Amount:      200.0,
		Description: "Test income",
		Category:    "salary",
		Source:      "employer",
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	mockIncomeService.On("CreateIncome", context.Background(), incomeRequest).Return(incomeResponse, nil)

	result, err = factory.CreateTransaction(context.Background(), "user1", 200.0, "Test income", "salary", nil)
	assert.NoError(t, err)
	assert.Equal(t, incomeResponse.ID, result.GetID())
	assert.Equal(t, incomeResponse.UserID, result.GetUserID())
	assert.Equal(t, incomeResponse.Amount, result.GetAmount())
	assert.Equal(t, incomeResponse.Description, result.GetDescription())
	assert.Equal(t, incomeResponse.Category, result.GetCategory())
	mockIncomeService.AssertExpectations(t)
}

func TestGetTransaction(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test obtener un gasto
	expenseResponse := &expenses.GetExpenseResponse{
		ID:          "1",
		UserID:      "user1",
		Amount:      100.0,
		Description: "Test expense",
		Category:    "food",
		Paid:        false,
		DueDate:     time.Now().Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	mockExpenseService.On("GetExpense", context.Background(), "user1", "1").Return(expenseResponse, nil)

	result, err := factory.GetTransaction(context.Background(), "user1", "1", ExpenseType)
	assert.NoError(t, err)
	assert.Equal(t, expenseResponse.ID, result.GetID())
	assert.Equal(t, expenseResponse.UserID, result.GetUserID())
	assert.Equal(t, expenseResponse.Amount, result.GetAmount())
	assert.Equal(t, expenseResponse.Description, result.GetDescription())
	assert.Equal(t, expenseResponse.Category, result.GetCategory())
	mockExpenseService.AssertExpectations(t)

	// Test obtener un ingreso
	incomeResponse := &incomes.GetIncomeResponse{
		ID:          "2",
		UserID:      "user1",
		Amount:      200.0,
		Description: "Test income",
		Category:    "salary",
		Source:      "employer",
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	mockIncomeService.On("GetIncome", context.Background(), "user1", "2").Return(incomeResponse, nil)

	result, err = factory.GetTransaction(context.Background(), "user1", "2", IncomeType)
	assert.NoError(t, err)
	assert.Equal(t, incomeResponse.ID, result.GetID())
	assert.Equal(t, incomeResponse.UserID, result.GetUserID())
	assert.Equal(t, incomeResponse.Amount, result.GetAmount())
	assert.Equal(t, incomeResponse.Description, result.GetDescription())
	assert.Equal(t, incomeResponse.Category, result.GetCategory())
	mockIncomeService.AssertExpectations(t)
}

func TestListTransactions(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test listar gastos
	expenseResponse := &expenses.ListExpensesResponse{
		Expenses: []expenses.GetExpenseResponse{
			{
				ID:          "1",
				UserID:      "user1",
				Amount:      100.0,
				Description: "Test expense 1",
				Category:    "food",
				Paid:        false,
				DueDate:     time.Now().Format("2006-01-02T15:04:05Z07:00"),
				CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			},
			{
				ID:          "2",
				UserID:      "user1",
				Amount:      200.0,
				Description: "Test expense 2",
				Category:    "transport",
				Paid:        true,
				DueDate:     time.Now().Format("2006-01-02T15:04:05Z07:00"),
				CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			},
		},
	}

	mockExpenseService.On("ListExpenses", context.Background(), "user1").Return(expenseResponse, nil)

	results, err := factory.ListTransactions(context.Background(), "user1", ExpenseType)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, expenseResponse.Expenses[0].ID, results[0].GetID())
	assert.Equal(t, expenseResponse.Expenses[1].ID, results[1].GetID())
	mockExpenseService.AssertExpectations(t)

	// Test listar ingresos
	incomeResponse := &incomes.ListIncomesResponse{
		Incomes: []incomes.GetIncomeResponse{
			{
				ID:          "3",
				UserID:      "user1",
				Amount:      300.0,
				Description: "Test income 1",
				Category:    "salary",
				Source:      "employer",
				CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			},
			{
				ID:          "4",
				UserID:      "user1",
				Amount:      400.0,
				Description: "Test income 2",
				Category:    "freelance",
				Source:      "client",
				CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			},
		},
	}

	mockIncomeService.On("ListIncomes", context.Background(), "user1").Return(incomeResponse, nil)

	results, err = factory.ListTransactions(context.Background(), "user1", IncomeType)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, incomeResponse.Incomes[0].ID, results[0].GetID())
	assert.Equal(t, incomeResponse.Incomes[1].ID, results[1].GetID())
	mockIncomeService.AssertExpectations(t)
}

func TestUpdateTransaction(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test actualizar un gasto
	dueDate := time.Now().AddDate(0, 1, 0)
	expenseRequest := &expenses.UpdateExpenseRequest{
		Amount:      150.0,
		Description: "Updated expense",
		Category:    "entertainment",
		Paid:        true,
		DueDate:     dueDate.Format("2006-01-02T15:04:05Z07:00"),
	}

	expenseResponse := &expenses.UpdateExpenseResponse{
		ID:          "1",
		UserID:      "user1",
		Amount:      150.0,
		Description: "Updated expense",
		Category:    "entertainment",
		Paid:        true,
		DueDate:     dueDate.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	mockExpenseService.On("UpdateExpense", context.Background(), "user1", "1", expenseRequest).Return(expenseResponse, nil)

	result, err := factory.UpdateTransaction(context.Background(), "user1", "1", 150.0, "Updated expense", "entertainment", &dueDate, ExpenseType)
	assert.NoError(t, err)
	assert.Equal(t, expenseResponse.ID, result.GetID())
	assert.Equal(t, expenseResponse.Amount, result.GetAmount())
	assert.Equal(t, expenseResponse.Description, result.GetDescription())
	assert.Equal(t, expenseResponse.Category, result.GetCategory())
	mockExpenseService.AssertExpectations(t)

	// Test actualizar un ingreso
	incomeRequest := &incomes.UpdateIncomeRequest{
		Amount:      250.0,
		Description: "Updated income",
		Category:    "bonus",
		Source:      "employer",
	}

	incomeResponse := &incomes.UpdateIncomeResponse{
		ID:          "2",
		UserID:      "user1",
		Amount:      250.0,
		Description: "Updated income",
		Category:    "bonus",
		Source:      "employer",
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	mockIncomeService.On("UpdateIncome", context.Background(), "user1", "2", incomeRequest).Return(incomeResponse, nil)

	result, err = factory.UpdateTransaction(context.Background(), "user1", "2", 250.0, "Updated income", "bonus", nil, IncomeType)
	assert.NoError(t, err)
	assert.Equal(t, incomeResponse.ID, result.GetID())
	assert.Equal(t, incomeResponse.Amount, result.GetAmount())
	assert.Equal(t, incomeResponse.Description, result.GetDescription())
	assert.Equal(t, incomeResponse.Category, result.GetCategory())
	mockIncomeService.AssertExpectations(t)
}

func TestDeleteTransaction(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test eliminar un gasto
	mockExpenseService.On("DeleteExpense", context.Background(), "user1", "1").Return(nil)

	err := factory.DeleteTransaction(context.Background(), "user1", "1", ExpenseType)
	assert.NoError(t, err)
	mockExpenseService.AssertExpectations(t)

	// Test eliminar un ingreso
	mockIncomeService.On("DeleteIncome", context.Background(), "user1", "2").Return(nil)

	err = factory.DeleteTransaction(context.Background(), "user1", "2", IncomeType)
	assert.NoError(t, err)
	mockIncomeService.AssertExpectations(t)
}

func TestCreateTransactionError(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test error al crear un gasto
	dueDate := time.Now().AddDate(0, 1, 0)
	expenseRequest := &expenses.CreateExpenseRequest{
		UserID:      "user1",
		Amount:      100.0,
		Description: "Test expense",
		Category:    "food",
		Paid:        false,
		DueDate:     dueDate.Format("2006-01-02T15:04:05Z07:00"),
	}

	expectedError := errors.New("error al crear gasto")
	mockExpenseService.On("CreateExpense", context.Background(), expenseRequest).Return(nil, expectedError)

	result, err := factory.CreateTransaction(context.Background(), "user1", 100.0, "Test expense", "food", &dueDate)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
	mockExpenseService.AssertExpectations(t)

	// Test error al crear un ingreso
	incomeRequest := &incomes.CreateIncomeRequest{
		UserID:      "user1",
		Amount:      200.0,
		Description: "Test income",
		Category:    "salary",
		Source:      "employer",
	}

	expectedError = errors.New("error al crear ingreso")
	mockIncomeService.On("CreateIncome", context.Background(), incomeRequest).Return(nil, expectedError)

	result, err = factory.CreateTransaction(context.Background(), "user1", 200.0, "Test income", "salary", nil)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
	mockIncomeService.AssertExpectations(t)
}

func TestGetTransactionError(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test error al obtener un gasto
	expectedError := errors.New("error al obtener gasto")
	mockExpenseService.On("GetExpense", context.Background(), "user1", "1").Return(nil, expectedError)

	result, err := factory.GetTransaction(context.Background(), "user1", "1", ExpenseType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
	mockExpenseService.AssertExpectations(t)

	// Test error al obtener un ingreso
	expectedError = errors.New("error al obtener ingreso")
	mockIncomeService.On("GetIncome", context.Background(), "user1", "2").Return(nil, expectedError)

	result, err = factory.GetTransaction(context.Background(), "user1", "2", IncomeType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
	mockIncomeService.AssertExpectations(t)

	// Test tipo de transacción inválido
	result, err = factory.GetTransaction(context.Background(), "user1", "3", "invalid")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidTransactionType, err)
	assert.Nil(t, result)
}

func TestListTransactionsError(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test error al listar gastos
	expectedError := errors.New("error al listar gastos")
	mockExpenseService.On("ListExpenses", context.Background(), "user1").Return(nil, expectedError)

	results, err := factory.ListTransactions(context.Background(), "user1", ExpenseType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, results)
	mockExpenseService.AssertExpectations(t)

	// Test error al listar ingresos
	expectedError = errors.New("error al listar ingresos")
	mockIncomeService.On("ListIncomes", context.Background(), "user1").Return(nil, expectedError)

	results, err = factory.ListTransactions(context.Background(), "user1", IncomeType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, results)
	mockIncomeService.AssertExpectations(t)

	// Test tipo de transacción inválido
	results, err = factory.ListTransactions(context.Background(), "user1", "invalid")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidTransactionType, err)
	assert.Nil(t, results)
}

func TestUpdateTransactionError(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test error al actualizar un gasto
	dueDate := time.Now().AddDate(0, 1, 0)
	expenseRequest := &expenses.UpdateExpenseRequest{
		Amount:      150.0,
		Description: "Updated expense",
		Category:    "entertainment",
		Paid:        true,
		DueDate:     dueDate.Format("2006-01-02T15:04:05Z07:00"),
	}

	expectedError := errors.New("error al actualizar gasto")
	mockExpenseService.On("UpdateExpense", context.Background(), "user1", "1", expenseRequest).Return(nil, expectedError)

	result, err := factory.UpdateTransaction(context.Background(), "user1", "1", 150.0, "Updated expense", "entertainment", &dueDate, ExpenseType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
	mockExpenseService.AssertExpectations(t)

	// Test error al actualizar un ingreso
	incomeRequest := &incomes.UpdateIncomeRequest{
		Amount:      250.0,
		Description: "Updated income",
		Category:    "bonus",
		Source:      "employer",
	}

	expectedError = errors.New("error al actualizar ingreso")
	mockIncomeService.On("UpdateIncome", context.Background(), "user1", "2", incomeRequest).Return(nil, expectedError)

	result, err = factory.UpdateTransaction(context.Background(), "user1", "2", 250.0, "Updated income", "bonus", nil, IncomeType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
	mockIncomeService.AssertExpectations(t)

	// Test tipo de transacción inválido
	result, err = factory.UpdateTransaction(context.Background(), "user1", "3", 300.0, "Invalid", "category", nil, "invalid")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidTransactionType, err)
	assert.Nil(t, result)
}

func TestDeleteTransactionError(t *testing.T) {
	mockIncomeService := new(MockIncomeService)
	mockExpenseService := new(MockExpenseService)
	factory := NewTransactionFactory(mockIncomeService, mockExpenseService)

	// Test error al eliminar un gasto
	expectedError := errors.New("error al eliminar gasto")
	mockExpenseService.On("DeleteExpense", context.Background(), "user1", "1").Return(expectedError)

	err := factory.DeleteTransaction(context.Background(), "user1", "1", ExpenseType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockExpenseService.AssertExpectations(t)

	// Test error al eliminar un ingreso
	expectedError = errors.New("error al eliminar ingreso")
	mockIncomeService.On("DeleteIncome", context.Background(), "user1", "2").Return(expectedError)

	err = factory.DeleteTransaction(context.Background(), "user1", "2", IncomeType)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockIncomeService.AssertExpectations(t)

	// Test tipo de transacción inválido
	err = factory.DeleteTransaction(context.Background(), "user1", "3", "invalid")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidTransactionType, err)
}

func TestHelperFunctions(t *testing.T) {
	// Test toTransaction con tipo inválido
	result := toTransaction("invalid")
	assert.Nil(t, result)

	// Test toTransactionSlice con tipo inválido
	results := toTransactionSlice("invalid")
	assert.Empty(t, results)

	// Test parseTime con formato inválido
	invalidTime := parseTime("invalid")
	assert.Equal(t, time.Time{}, invalidTime)
}
