package categories

import (
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(categoryID string, category *CategoryModel) error {
	category.ID = categoryID
	result := r.db.Create(category)
	return result.Error
}

func (r *CategoryRepository) List() ([]CategoryModel, error) {
	var categories []CategoryModel
	result := r.db.Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) Get(id string) (*CategoryModel, error) {
	var category CategoryModel
	result := r.db.First(&category, id)
	return &category, result.Error
}

func (r *CategoryRepository) Update(category *CategoryModel) error {
	result := r.db.Save(category)
	return result.Error
}

func (r *CategoryRepository) Delete(id string) error {
	result := r.db.Delete(&CategoryModel{}, id)
	return result.Error
}
