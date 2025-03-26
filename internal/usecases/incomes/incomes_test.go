package incomes

import (
	"context"
	"testing"

	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockIncomeRepository es un mock del repositorio de ingresos
type MockIncomeRepository struct {
	mock.Mock
}

func (m *MockIncomeRepository) Create(income *domain.Income) error {
	args := m.Called(income)
	return args.Error(0)
}

func (m *MockIncomeRepository) Get(userID string, incomeID string) (*domain.Income, error) {
	args := m.Called(userID, incomeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Income), args.Error(1)
}

func (m *MockIncomeRepository) List(userID string) ([]*domain.Income, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Income), args.Error(1)
}

func (m *MockIncomeRepository) Update(income *domain.Income) error {
	args := m.Called(income)
	return args.Error(0)
}

func (m *MockIncomeRepository) Delete(userID string, incomeID string) error {
	args := m.Called(userID, incomeID)
	return args.Error(0)
}

func TestCreateIncome(t *testing.T) {
	mockRepo := new(MockIncomeRepository)
	service := NewIncomeService(mockRepo)

	request := &CreateIncomeRequest{
		UserID:      "user1",
		Amount:      1000.0,
		Description: "Test income",
		Category:    "salary",
		Source:      "employer",
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.Income")).Return(nil)

	result, err := service.CreateIncome(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, request.UserID, result.UserID)
	assert.Equal(t, request.Amount, result.Amount)
	assert.Equal(t, request.Description, result.Description)
	assert.Equal(t, request.Category, result.Category)
	assert.Equal(t, request.Source, result.Source)
	mockRepo.AssertExpectations(t)
}

func TestGetIncome(t *testing.T) {
	mockRepo := new(MockIncomeRepository)
	service := NewIncomeService(mockRepo)

	expectedIncome := domain.NewIncomeBuilder().
		SetID("1").
		SetUserID("user1").
		SetAmount(1000.0).
		SetDescription("Test income").
		SetCategory("salary").
		SetSource("employer").
		Build()

	mockRepo.On("Get", "user1", "1").Return(expectedIncome, nil)

	result, err := service.GetIncome(context.Background(), "user1", "1")
	assert.NoError(t, err)
	assert.Equal(t, expectedIncome.ID, result.ID)
	assert.Equal(t, expectedIncome.UserID, result.UserID)
	assert.Equal(t, expectedIncome.Amount, result.Amount)
	assert.Equal(t, expectedIncome.Description, result.Description)
	assert.Equal(t, expectedIncome.Category, result.Category)
	assert.Equal(t, expectedIncome.Source, result.Source)
	mockRepo.AssertExpectations(t)
}

func TestListIncomes(t *testing.T) {
	mockRepo := new(MockIncomeRepository)
	service := NewIncomeService(mockRepo)

	income1 := domain.NewIncomeBuilder().
		SetID("1").
		SetUserID("user1").
		SetAmount(1000.0).
		SetDescription("Income 1").
		SetCategory("salary").
		SetSource("employer").
		Build()

	income2 := domain.NewIncomeBuilder().
		SetID("2").
		SetUserID("user1").
		SetAmount(500.0).
		SetDescription("Income 2").
		SetCategory("freelance").
		SetSource("client").
		Build()

	expectedIncomes := []*domain.Income{income1, income2}

	mockRepo.On("List", "user1").Return(expectedIncomes, nil)

	result, err := service.ListIncomes(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Len(t, result.Incomes, 2)
	assert.Equal(t, income1.ID, result.Incomes[0].ID)
	assert.Equal(t, income2.ID, result.Incomes[1].ID)
	mockRepo.AssertExpectations(t)
}

func TestUpdateIncome(t *testing.T) {
	mockRepo := new(MockIncomeRepository)
	service := NewIncomeService(mockRepo)

	existingIncome := domain.NewIncomeBuilder().
		SetID("1").
		SetUserID("user1").
		SetAmount(1000.0).
		SetDescription("Test income").
		SetCategory("salary").
		SetSource("employer").
		Build()

	request := &UpdateIncomeRequest{
		Amount:      1500.0,
		Description: "Updated income",
		Category:    "bonus",
		Source:      "employer",
	}

	updatedIncome := domain.NewIncomeBuilder().
		SetID("1").
		SetUserID("user1").
		SetAmount(request.Amount).
		SetDescription(request.Description).
		SetCategory(request.Category).
		SetSource(request.Source).
		Build()

	mockRepo.On("Get", "user1", "1").Return(existingIncome, nil)
	mockRepo.On("Update", mock.AnythingOfType("*domain.Income")).Return(nil)

	result, err := service.UpdateIncome(context.Background(), "user1", "1", request)
	assert.NoError(t, err)
	assert.Equal(t, updatedIncome.ID, result.ID)
	assert.Equal(t, request.Amount, result.Amount)
	assert.Equal(t, request.Description, result.Description)
	assert.Equal(t, request.Category, result.Category)
	assert.Equal(t, request.Source, result.Source)
	mockRepo.AssertExpectations(t)
}

func TestDeleteIncome(t *testing.T) {
	mockRepo := new(MockIncomeRepository)
	service := NewIncomeService(mockRepo)

	income := domain.NewIncomeBuilder().
		SetID("1").
		SetUserID("user1").
		SetAmount(1000.0).
		SetDescription("Test income").
		SetCategory("salary").
		SetSource("employer").
		Build()

	mockRepo.On("Get", "user1", "1").Return(income, nil)
	mockRepo.On("Delete", "user1", "1").Return(nil)

	err := service.DeleteIncome(context.Background(), "user1", "1")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
