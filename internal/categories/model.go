package categories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryModel struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
}

func (c *CategoryModel) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
