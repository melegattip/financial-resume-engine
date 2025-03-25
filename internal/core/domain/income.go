package domain

import "time"

// Income representa un ingreso
type Income struct {
	BaseTransaction
	Source string `json:"source"`
}

// IncomeFactory implementa TransactionFactory para crear ingresos
type IncomeFactory struct{}

func NewIncomeFactory() *IncomeFactory {
	return &IncomeFactory{}
}

func (f *IncomeFactory) CreateTransaction() Transaction {
	return &Income{
		BaseTransaction: BaseTransaction{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// Builder para Income
type IncomeBuilder struct {
	income *Income
}

func NewIncomeBuilder() *IncomeBuilder {
	return &IncomeBuilder{
		income: &Income{
			BaseTransaction: BaseTransaction{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}
}

func (b *IncomeBuilder) SetID(id string) *IncomeBuilder {
	b.income.ID = id
	return b
}

func (b *IncomeBuilder) SetUserID(userID string) *IncomeBuilder {
	b.income.UserID = userID
	return b
}

func (b *IncomeBuilder) SetAmount(amount float64) *IncomeBuilder {
	b.income.Amount = amount
	return b
}

func (b *IncomeBuilder) SetDescription(description string) *IncomeBuilder {
	b.income.Description = description
	return b
}

func (b *IncomeBuilder) SetCategory(category string) *IncomeBuilder {
	b.income.Category = category
	return b
}

func (b *IncomeBuilder) SetSource(source string) *IncomeBuilder {
	b.income.Source = source
	return b
}

func (b *IncomeBuilder) SetCreatedAt(createdAt time.Time) *IncomeBuilder {
	b.income.CreatedAt = createdAt
	return b
}

func (b *IncomeBuilder) SetUpdatedAt(updatedAt time.Time) *IncomeBuilder {
	b.income.UpdatedAt = updatedAt
	return b
}

func (b *IncomeBuilder) Build() *Income {
	return b.income
}
