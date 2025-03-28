package categories

import (
	"testing"

	"github.com/google/uuid"
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
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

func (m *MockCategoryRepository) Get(name string) (*domain.Category, error) {
	args := m.Called(name)
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
	tests := []struct {
		name          string
		category      *domain.Category
		mockSetup     func(*MockCategoryRepository)
		expectedError error
	}{
		{
			name: "create category successfully",
			category: domain.NewCategoryBuilder().
				SetName("Test Category").
				Build(),
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Get", "Test Category").Return(nil, errors.NewResourceNotFound(""))
				m.On("Create", mock.AnythingOfType("*domain.Category")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "error when creating category that already exists",
			category: domain.NewCategoryBuilder().
				SetName("Existing Category").
				Build(),
			mockSetup: func(m *MockCategoryRepository) {
				existingCategory := domain.NewCategoryBuilder().
					SetID("cat_" + uuid.New().String()[:8]).
					SetName("Existing Category").
					Build()
				m.On("Get", "Existing Category").Return(existingCategory, nil)
			},
			expectedError: errors.NewResourceAlreadyExists("Error creating category"),
		},
		{
			name: "error when creating category with empty name",
			category: &domain.Category{
				Name: "",
			},
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Get", "").Return(nil, errors.NewResourceNotFound(""))
			},
			expectedError: errors.NewBadRequest("Category name cannot be empty"),
		},
		{
			name:          "error when creating nil category",
			category:      nil,
			mockSetup:     func(m *MockCategoryRepository) {},
			expectedError: errors.NewBadRequest("Category cannot be nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCategoryRepository)
			service := NewCreateCategory(mockRepo)
			tt.mockSetup(mockRepo)

			result, err := service.Execute(tt.category)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.IsType(t, tt.expectedError, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, result.ID)
			assert.Equal(t, tt.category.Name, result.Name)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetCategory(t *testing.T) {
	tests := []struct {
		name          string
		categoryName  string
		mockSetup     func(*MockCategoryRepository)
		expectedError error
	}{
		{
			name:         "obtener categoría existente",
			categoryName: "Test Category",
			mockSetup: func(m *MockCategoryRepository) {
				expectedCategory := domain.NewCategoryBuilder().
					SetID("cat_" + uuid.New().String()[:8]).
					SetName("Test Category").
					Build()
				m.On("Get", "Test Category").Return(expectedCategory, nil)
			},
			expectedError: nil,
		},
		{
			name:         "error al obtener categoría que no existe",
			categoryName: "Non Existent",
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Get", "Non Existent").Return(nil, errors.NewResourceNotFound(""))
			},
			expectedError: errors.NewResourceNotFound(""),
		},
		{
			name:         "error al obtener categoría con nombre vacío",
			categoryName: "",
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Get", "").Return(nil, errors.NewBadRequest(""))
			},
			expectedError: errors.NewBadRequest(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCategoryRepository)
			service := NewGetCategory(mockRepo)
			tt.mockSetup(mockRepo)

			category, err := service.Execute(tt.categoryName)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, category)
				assert.IsType(t, tt.expectedError, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, category)
			assert.Equal(t, tt.categoryName, category.Name)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestListCategories(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(*MockCategoryRepository)
		expectedError error
	}{
		{
			name: "listar categorías exitosamente",
			mockSetup: func(m *MockCategoryRepository) {
				category1 := domain.NewCategoryBuilder().
					SetID("cat_" + uuid.New().String()[:8]).
					SetName("Category 1").
					Build()

				category2 := domain.NewCategoryBuilder().
					SetID("cat_" + uuid.New().String()[:8]).
					SetName("Category 2").
					Build()

				expectedCategories := []*domain.Category{category1, category2}
				m.On("List").Return(expectedCategories, nil)
			},
			expectedError: nil,
		},
		{
			name: "listar categorías vacías",
			mockSetup: func(m *MockCategoryRepository) {
				m.On("List").Return([]*domain.Category{}, nil)
			},
			expectedError: nil,
		},
		{
			name: "error al listar categorías",
			mockSetup: func(m *MockCategoryRepository) {
				m.On("List").Return(nil, errors.NewInternalServerError(""))
			},
			expectedError: errors.NewInternalServerError(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCategoryRepository)
			service := NewListCategories(mockRepo)
			tt.mockSetup(mockRepo)

			categories, err := service.Execute()
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, categories)
				assert.IsType(t, tt.expectedError, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, categories)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		name          string
		category      *domain.Category
		mockSetup     func(*MockCategoryRepository)
		expectedError error
	}{
		{
			name: "update category successfully",
			category: domain.NewCategoryBuilder().
				SetID("cat_" + uuid.New().String()[:8]).
				SetName("Updated Category").
				Build(),
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Update", mock.AnythingOfType("*domain.Category")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "error when updating non-existent category",
			category: domain.NewCategoryBuilder().
				SetID("cat_" + uuid.New().String()[:8]).
				SetName("Non Existent").
				Build(),
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Update", mock.AnythingOfType("*domain.Category")).Return(errors.NewResourceNotFound(""))
			},
			expectedError: errors.NewResourceNotFound(""),
		},
		{
			name: "error when updating category with invalid ID",
			category: domain.NewCategoryBuilder().
				SetID("invalid_id").
				SetName("Invalid Category").
				Build(),
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Update", mock.AnythingOfType("*domain.Category")).Return(errors.NewBadRequest(""))
			},
			expectedError: errors.NewBadRequest("Invalid category ID"),
		},
		{
			name: "error when updating category with empty name",
			category: &domain.Category{
				ID:   "cat_" + uuid.New().String()[:8],
				Name: "",
			},
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Update", mock.AnythingOfType("*domain.Category")).Return(errors.NewBadRequest(""))
			},
			expectedError: errors.NewBadRequest("Category name cannot be empty"),
		},
		{
			name:          "error when updating nil category",
			category:      nil,
			mockSetup:     func(m *MockCategoryRepository) {},
			expectedError: errors.NewBadRequest("Category cannot be nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCategoryRepository)
			service := NewUpdateCategory(mockRepo)
			tt.mockSetup(mockRepo)

			err := service.Execute(tt.category)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				return
			}

			assert.NoError(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		name          string
		categoryID    string
		mockSetup     func(*MockCategoryRepository)
		expectedError error
	}{
		{
			name:       "eliminar categoría exitosamente",
			categoryID: "cat_" + uuid.New().String()[:8],
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Delete", mock.AnythingOfType("string")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:       "error al eliminar categoría que no existe",
			categoryID: "cat_" + uuid.New().String()[:8],
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Delete", mock.AnythingOfType("string")).Return(errors.NewResourceNotFound(""))
			},
			expectedError: errors.NewResourceNotFound(""),
		},
		{
			name:       "error al eliminar categoría con ID inválido",
			categoryID: "invalid_id",
			mockSetup: func(m *MockCategoryRepository) {
				m.On("Delete", "invalid_id").Return(errors.NewBadRequest(""))
			},
			expectedError: errors.NewBadRequest(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockCategoryRepository)
			service := NewDeleteCategory(mockRepo)
			tt.mockSetup(mockRepo)

			err := service.Execute(tt.categoryID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
				return
			}

			assert.NoError(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
