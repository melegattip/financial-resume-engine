package transactions

import (
	"time"
)

type TransactionModel struct {
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	TypeID      string    `gorm:"column:type_id" json:"type_id"`
	Description string    `gorm:"column:description" json:"description"`
	Amount      float64   `gorm:"column:amount" json:"amount"`
	Payed       bool      `gorm:"column:payed" json:"payed"`
	ExpiryDate  time.Time `gorm:"column:expiry_date" json:"expiry_date"`
	CategoryID  string    `gorm:"column:category" json:"category"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (TransactionModel) TableName() string {
	return "transactions"
}

// Builder pattern
type TransactionBuilder struct {
	transaction TransactionModel
}

func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{}
}

func (b *TransactionBuilder) SetUserID(userID string) *TransactionBuilder {
	b.transaction.UserID = userID
	return b
}

func (b *TransactionBuilder) SetID(id string) *TransactionBuilder {
	b.transaction.ID = id
	return b
}

func (b *TransactionBuilder) SetTypeID(typeID string) *TransactionBuilder {
	b.transaction.TypeID = typeID
	return b
}

func (b *TransactionBuilder) SetDescription(description string) *TransactionBuilder {
	b.transaction.Description = description
	return b
}

func (b *TransactionBuilder) SetAmount(amount float64) *TransactionBuilder {
	b.transaction.Amount = amount
	return b
}

func (b *TransactionBuilder) SetPayed(payed bool) *TransactionBuilder {
	b.transaction.Payed = payed
	return b
}

func (b *TransactionBuilder) SetExpiryDate(expiryDate time.Time) *TransactionBuilder {
	b.transaction.ExpiryDate = expiryDate
	return b
}

func (b *TransactionBuilder) SetCategoryID(categoryID string) *TransactionBuilder {
	b.transaction.CategoryID = categoryID
	return b
}

func (b *TransactionBuilder) SetCreatedAt(createdAt time.Time) *TransactionBuilder {
	b.transaction.CreatedAt = createdAt
	return b
}

func (b *TransactionBuilder) SetUpdatedAt(updatedAt time.Time) *TransactionBuilder {
	b.transaction.UpdatedAt = updatedAt
	return b
}

func (b *TransactionBuilder) Build() *TransactionModel {
	return &b.transaction
}
