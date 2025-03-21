package categories

import (
	"time"
)

type CategoryModel struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	UserID      string    `gorm:"column:user_id;not null" json:"user_id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (CategoryModel) TableName() string {
	return "categories"
}

// Builder pattern
type CategoryBuilder struct {
	category CategoryModel
}

func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{}
}

func (b *CategoryBuilder) SetID(id string) *CategoryBuilder {
	b.category.ID = id
	return b
}

func (b *CategoryBuilder) SetUserID(userID string) *CategoryBuilder {
	b.category.UserID = userID
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

func (b *CategoryBuilder) Build() *CategoryModel {
	return &b.category
}
