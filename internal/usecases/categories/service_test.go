package categories

import (
	"testing"

	"github.com/google/uuid"
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository es un mock del repositorio de categorías
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(category *domain.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Get(id string) (*domain.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) List() ([]*domain.Category, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(category *domain.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCategory(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewCreateCategory(mockRepo)

	category := domain.NewCategoryBuilder().
		SetName("Test Category").
		Build()

	mockRepo.On("Create", mock.AnythingOfType("*domain.Category")).Return(nil)

	result, err := service.Execute(category)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetCategory(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewGetCategory(mockRepo)

	expectedCategory := domain.NewCategoryBuilder().
		SetID("cat_" + uuid.New().String()[:8]).
		SetName("Test Category").
		Build()

	mockRepo.On("Get", "1").Return(expectedCategory, nil)

	category, err := service.Execute("1")
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
	mockRepo.AssertExpectations(t)
}

func TestListCategories(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewListCategories(mockRepo)

	category1 := domain.NewCategoryBuilder().
		SetID("cat_" + uuid.New().String()[:8]).
		SetName("Category 1").
		Build()

	category2 := domain.NewCategoryBuilder().
		SetID("cat_" + uuid.New().String()[:8]).
		SetName("Category 2").
		Build()

	expectedCategories := []*domain.Category{category1, category2}

	mockRepo.On("List").Return(expectedCategories, nil)

	categories, err := service.Execute()
	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewUpdateCategory(mockRepo)

	category := domain.NewCategoryBuilder().
		SetID("cat_" + uuid.New().String()[:8]).
		SetName("Updated Category").
		Build()

	mockRepo.On("Update", category).Return(nil)

	err := service.Execute(category)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewDeleteCategory(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := service.Execute("1")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
