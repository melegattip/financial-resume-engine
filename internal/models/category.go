package models

import (
	"time"
)

type CategoryModel struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
