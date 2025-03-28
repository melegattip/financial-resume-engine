package repository

import (
	"errors"

	"github.com/melegattip/financial-resume-engine/internal/core/domain"
	coreErrors "github.com/melegattip/financial-resume-engine/internal/core/errors"
	"gorm.io/gorm"
)

// Category implementa el repositorio de categorías
type Category struct {
	db *gorm.DB
}

// NewCategoryRepository crea una nueva instancia del repositorio de categorías
func NewCategoryRepository(db *gorm.DB) *Category {
	return &Category{db: db}
}

func (r *Category) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *Category) Get(name string) (*domain.Category, error) {
	var category domain.Category
	result := r.db.First(&category, "name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, coreErrors.NewResourceNotFound("Category not found")
		}
		return nil, result.Error
	}
	return &category, nil
}

func (r *Category) List() ([]*domain.Category, error) {
	var categories []domain.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	categoryPointers := make([]*domain.Category, len(categories))
	for i := range categories {
		categoryPointers[i] = &categories[i]
	}
	return categoryPointers, nil
}

func (r *Category) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *Category) Delete(id string) error {
	result := r.db.Delete(&domain.Category{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return coreErrors.NewResourceNotFound("Category not found")
	}
	return nil
}
