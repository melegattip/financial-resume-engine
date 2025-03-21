package categories

import (
	"time"

	"github.com/google/uuid"
)

type CreateCategory struct {
	CategoryRepository *CategoryRepository
}

func NewCreateCategory(repo *CategoryRepository) *CreateCategory {
	return &CreateCategory{CategoryRepository: repo}
}

func (s *CreateCategory) Execute(category *CategoryModel) (*CategoryModel, error) {
	// Generar ID único para la categoría
	categoryID := "cat_" + uuid.New().String()[:8]

	// Establecer fecha de creación
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = category.CreatedAt

	// Crear la categoría en el repositorio
	err := s.CategoryRepository.Create(categoryID, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

type ListCategories struct {
	CategoryRepository *CategoryRepository
}

func NewListCategories(repo *CategoryRepository) *ListCategories {
	return &ListCategories{CategoryRepository: repo}
}

func (s *ListCategories) Execute() ([]CategoryModel, error) {
	return s.CategoryRepository.List()
}

type GetCategory struct {
	CategoryRepository *CategoryRepository
}

func NewGetCategory(repo *CategoryRepository) *GetCategory {
	return &GetCategory{CategoryRepository: repo}
}

func (s *GetCategory) Execute(id string) (*CategoryModel, error) {
	return s.CategoryRepository.Get(id)
}

type UpdateCategory struct {
	CategoryRepository *CategoryRepository
}

func NewUpdateCategory(repo *CategoryRepository) *UpdateCategory {
	return &UpdateCategory{CategoryRepository: repo}
}

func (s *UpdateCategory) Execute(category *CategoryModel) error {
	category.UpdatedAt = time.Now().UTC()
	return s.CategoryRepository.Update(category)
}

type DeleteCategory struct {
	CategoryRepository *CategoryRepository
}

func NewDeleteCategory(repo *CategoryRepository) *DeleteCategory {
	return &DeleteCategory{CategoryRepository: repo}
}

func (s *DeleteCategory) Execute(id string) error {
	return s.CategoryRepository.Delete(id)
}
