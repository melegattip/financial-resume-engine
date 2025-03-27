package categories

import (
	"strings"

	"github.com/google/uuid"
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/melegattip/financial-resume-engine/internal/core/logs"
	baseRepo "github.com/melegattip/financial-resume-engine/internal/core/repository"
)

type CreateCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewCreateCategory(repo baseRepo.CategoryRepository) *CreateCategory {
	return &CreateCategory{CategoryRepository: repo}
}

func (s *CreateCategory) Execute(category *domain.Category) (*domain.Category, error) {
	// Validate category is not nil
	if category == nil {
		return nil, errors.NewBadRequest("Category cannot be nil")
	}

	// Validate empty name
	if strings.TrimSpace(category.Name) == "" {
		return nil, errors.NewBadRequest("Category name cannot be empty")
	}

	// Validate if category already exists
	existingCategory, err := s.CategoryRepository.Get(category.Name)
	if err != nil && !errors.IsResourceNotFound(err) {
		return nil, err
	}
	if existingCategory != nil {
		return nil, errors.NewResourceAlreadyExists(logs.ErrorCreatingCategory.GetMessage())
	}

	// Generate unique ID for category
	categoryID := "cat_" + uuid.New().String()[:8]

	// Set creation date
	category.ID = categoryID

	// Create category in repository
	err = s.CategoryRepository.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

type ListCategories struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewListCategories(repo baseRepo.CategoryRepository) *ListCategories {
	return &ListCategories{CategoryRepository: repo}
}

func (s *ListCategories) Execute() ([]*domain.Category, error) {
	return s.CategoryRepository.List()
}

type GetCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewGetCategory(repo baseRepo.CategoryRepository) *GetCategory {
	return &GetCategory{CategoryRepository: repo}
}

func (s *GetCategory) Execute(name string) (*domain.Category, error) {
	// Validate empty name
	if strings.TrimSpace(name) == "" {
		return nil, errors.NewBadRequest("Category name cannot be empty")
	}

	return s.CategoryRepository.Get(name)
}

type UpdateCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewUpdateCategory(repo baseRepo.CategoryRepository) *UpdateCategory {
	return &UpdateCategory{CategoryRepository: repo}
}

func (s *UpdateCategory) Execute(category *domain.Category) error {
	// Validate category is not nil
	if category == nil {
		return errors.NewBadRequest("Category cannot be nil")
	}

	// Validate empty name
	if strings.TrimSpace(category.Name) == "" {
		return errors.NewBadRequest("Category name cannot be empty")
	}

	// Validate invalid ID
	if !strings.HasPrefix(category.ID, "cat_") {
		return errors.NewBadRequest("Invalid category ID")
	}

	return s.CategoryRepository.Update(category)
}

type DeleteCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewDeleteCategory(repo baseRepo.CategoryRepository) *DeleteCategory {
	return &DeleteCategory{CategoryRepository: repo}
}

func (s *DeleteCategory) Execute(id string) error {
	// Validate invalid ID
	if !strings.HasPrefix(id, "cat_") {
		return errors.NewBadRequest("Invalid category ID")
	}

	return s.CategoryRepository.Delete(id)
}
