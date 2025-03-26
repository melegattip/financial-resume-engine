// Package domain defines the core business entities and their behavior
package domain

import (
	"errors"
	"time"
)

var (
	// ErrEmptyCategoryName is returned when attempting to create a category with an empty name
	ErrEmptyCategoryName = errors.New("category name cannot be empty")
)

// Category represents a transaction category in the system
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryBuilder implements the builder pattern for creating Category instances
type CategoryBuilder struct {
	category *Category
}

// NewCategoryBuilder creates a new CategoryBuilder instance with default timestamps
func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{
		category: &Category{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// SetID sets the category ID and returns the builder
func (b *CategoryBuilder) SetID(id string) *CategoryBuilder {
	b.category.ID = id
	return b
}

// SetName sets the category name and returns the builder
func (b *CategoryBuilder) SetName(name string) *CategoryBuilder {
	b.category.Name = name
	return b
}

// SetDescription sets the category description and returns the builder
func (b *CategoryBuilder) SetDescription(description string) *CategoryBuilder {
	b.category.Description = description
	return b
}

// SetCreatedAt sets the creation timestamp and returns the builder
func (b *CategoryBuilder) SetCreatedAt(createdAt time.Time) *CategoryBuilder {
	b.category.CreatedAt = createdAt
	return b
}

// SetUpdatedAt sets the last update timestamp and returns the builder
func (b *CategoryBuilder) SetUpdatedAt(updatedAt time.Time) *CategoryBuilder {
	b.category.UpdatedAt = updatedAt
	return b
}

// Build creates and returns a new Category instance if valid, nil otherwise
func (b *CategoryBuilder) Build() *Category {
	if b.category.Name == "" {
		return nil
	}
	return b.category
}

// Validate checks if the category is valid
// Returns ErrEmptyCategoryName if the name is empty
func (c *Category) Validate() error {
	if c.Name == "" {
		return ErrEmptyCategoryName
	}
	return nil
}
