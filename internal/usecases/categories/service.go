package categories

import (
	"time"

	"github.com/google/uuid"
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	baseRepo "github.com/melegattip/financial-resume-engine/internal/core/repository"
)

type CreateCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewCreateCategory(repo baseRepo.CategoryRepository) *CreateCategory {
	return &CreateCategory{CategoryRepository: repo}
}

func (s *CreateCategory) Execute(category *domain.Category) (*domain.Category, error) {
	// Generar ID único para la categoría
	categoryID := "cat_" + uuid.New().String()[:8]

	// Establecer fecha de creación
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = category.CreatedAt
	category.ID = categoryID

	// Crear la categoría en el repositorio
	err := s.CategoryRepository.Create(category)
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

func (s *GetCategory) Execute(id string) (*domain.Category, error) {
	return s.CategoryRepository.Get(id)
}

type UpdateCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewUpdateCategory(repo baseRepo.CategoryRepository) *UpdateCategory {
	return &UpdateCategory{CategoryRepository: repo}
}

func (s *UpdateCategory) Execute(category *domain.Category) error {
	category.UpdatedAt = time.Now().UTC()
	return s.CategoryRepository.Update(category)
}

type DeleteCategory struct {
	CategoryRepository baseRepo.CategoryRepository
}

func NewDeleteCategory(repo baseRepo.CategoryRepository) *DeleteCategory {
	return &DeleteCategory{CategoryRepository: repo}
}

func (s *DeleteCategory) Execute(id string) error {
	return s.CategoryRepository.Delete(id)
}
