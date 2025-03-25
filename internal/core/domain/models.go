package domain

import "time"

// BaseModel representa la estructura base para todos los modelos
type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User representa el modelo de usuario
type User struct {
	BaseModel
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Active    bool   `json:"active" gorm:"default:true"`
}

// FinancialData representa los datos financieros
type FinancialData struct {
	BaseModel
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Date        string  `json:"date"`
}
