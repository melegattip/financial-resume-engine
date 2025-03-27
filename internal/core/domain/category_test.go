package domain

import (
	"testing"

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
				Build()

			if tt.wantErr {
				assert.Nil(t, category)
				return
			}
			assert.NotNil(t, category)
			assert.Equal(t, tt.id, category.ID)
			assert.Equal(t, tt.title, category.Name)
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
				ID:   "123",
				Name: "Groceries",
			},
			wantErr: false,
		},
		{
			name: "Invalid - empty title",
			category: &Category{
				ID:   "123",
				Name: "",
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
				Build()

			assert.Equal(t, tt.id, category.ID)
			assert.Equal(t, tt.categoryName, category.Name)
		})
	}
}

func TestCategoryBuilder_CustomDates(t *testing.T) {
	category := NewCategoryBuilder().
		SetID("789").
		SetName("Entertainment").
		Build()

	assert.Equal(t, "789", category.ID)
	assert.Equal(t, "Entertainment", category.Name)
}
