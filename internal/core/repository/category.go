package repository

import (
	"github.com/melegattip/financial-resume-engine/internal/core/domain"
)

// CategoryRepository define las operaciones para el repositorio de categorías
type CategoryRepository interface {
	Create(category *domain.Category) error
	Get(id string) (*domain.Category, error)
	List() ([]*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id string) error
}
