package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		title       string
		description string
		wantErr     bool
	}{
		{
			name:        "Valid category",
			id:          "123",
			title:       "Groceries",
			description: "Food and household items",
			wantErr:     false,
		},
		{
			name:        "Empty title",
			id:          "123",
			title:       "",
			description: "Food and household items",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category := NewCategoryBuilder().
				SetID(tt.id).
				SetName(tt.title).
				SetDescription(tt.description).
				Build()

			if tt.wantErr {
				assert.Nil(t, category)
				return
			}
			assert.NotNil(t, category)
			assert.Equal(t, tt.id, category.ID)
			assert.Equal(t, tt.title, category.Name)
			assert.Equal(t, tt.description, category.Description)
		})
	}
}

func TestCategory_Validate(t *testing.T) {
	tests := []struct {
		name     string
		category *Category
		wantErr  bool
	}{
		{
			name: "Valid category",
			category: &Category{
				ID:          "123",
				Name:        "Groceries",
				Description: "Food and household items",
			},
			wantErr: false,
		},
		{
			name: "Invalid - empty title",
			category: &Category{
				ID:          "123",
				Name:        "",
				Description: "Food and household items",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.category.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ErrEmptyCategoryName, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestCategoryBuilder(t *testing.T) {
	tests := []struct {
		name         string
		id           string
		categoryName string
		description  string
	}{
		{
			name:         "Valid category",
			id:           "123",
			categoryName: "Groceries",
			description:  "Food and household items",
		},
		{
			name:         "Empty description",
			id:           "456",
			categoryName: "Transportation",
			description:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category := NewCategoryBuilder().
				SetID(tt.id).
				SetName(tt.categoryName).
				SetDescription(tt.description).
				Build()

			assert.Equal(t, tt.id, category.ID)
			assert.Equal(t, tt.categoryName, category.Name)
			assert.Equal(t, tt.description, category.Description)
			assert.False(t, category.CreatedAt.IsZero())
			assert.False(t, category.UpdatedAt.IsZero())
		})
	}
}

func TestCategoryBuilder_CustomDates(t *testing.T) {
	now := time.Now()
	category := NewCategoryBuilder().
		SetID("789").
		SetName("Entertainment").
		SetDescription("Movies and shows").
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Build()

	assert.Equal(t, "789", category.ID)
	assert.Equal(t, "Entertainment", category.Name)
	assert.Equal(t, "Movies and shows", category.Description)
	assert.Equal(t, now, category.CreatedAt)
	assert.Equal(t, now, category.UpdatedAt)
}
