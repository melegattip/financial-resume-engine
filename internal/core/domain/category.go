package domain

import (
	"errors"
	"time"
)

var (
	ErrEmptyCategoryName = errors.New("category name cannot be empty")
)

// Category representa una categoría de transacción
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryBuilder implementa el patrón builder para Category
type CategoryBuilder struct {
	category *Category
}

func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{
		category: &Category{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (b *CategoryBuilder) SetID(id string) *CategoryBuilder {
	b.category.ID = id
	return b
}

func (b *CategoryBuilder) SetName(name string) *CategoryBuilder {
	b.category.Name = name
	return b
}

func (b *CategoryBuilder) SetDescription(description string) *CategoryBuilder {
	b.category.Description = description
	return b
}

func (b *CategoryBuilder) SetCreatedAt(createdAt time.Time) *CategoryBuilder {
	b.category.CreatedAt = createdAt
	return b
}

func (b *CategoryBuilder) SetUpdatedAt(updatedAt time.Time) *CategoryBuilder {
	b.category.UpdatedAt = updatedAt
	return b
}

func (b *CategoryBuilder) Build() *Category {
	if b.category.Name == "" {
		return nil
	}
	return b.category
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return ErrEmptyCategoryName
	}
	return nil
}
