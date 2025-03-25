package domain

import "time"

// Expense representa un gasto
type Expense struct {
	BaseTransaction
	TypeID  int       `json:"type_id"`
	Paid    bool      `json:"paid,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
}

// ExpenseFactory implementa TransactionFactory para crear gastos
type ExpenseFactory struct{}

func NewExpenseFactory() *ExpenseFactory {
	return &ExpenseFactory{}
}

func (f *ExpenseFactory) CreateTransaction() Transaction {
	return &Expense{
		BaseTransaction: BaseTransaction{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		TypeID: 2, // 2 representa "expense"
	}
}

// Builder para Expense
type ExpenseBuilder struct {
	expense *Expense
}

func NewExpenseBuilder() *ExpenseBuilder {
	return &ExpenseBuilder{
		expense: &Expense{
			BaseTransaction: BaseTransaction{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			TypeID: 2, // 2 representa "expense"
		},
	}
}

func (b *ExpenseBuilder) SetID(id string) *ExpenseBuilder {
	b.expense.ID = id
	return b
}

func (b *ExpenseBuilder) SetUserID(userID string) *ExpenseBuilder {
	b.expense.UserID = userID
	return b
}

func (b *ExpenseBuilder) SetAmount(amount float64) *ExpenseBuilder {
	b.expense.Amount = amount
	return b
}

func (b *ExpenseBuilder) SetDescription(description string) *ExpenseBuilder {
	b.expense.Description = description
	return b
}

func (b *ExpenseBuilder) SetCategory(category string) *ExpenseBuilder {
	b.expense.Category = category
	return b
}

func (b *ExpenseBuilder) SetPaid(paid bool) *ExpenseBuilder {
	b.expense.Paid = paid
	return b
}

func (b *ExpenseBuilder) SetDueDate(dueDate time.Time) *ExpenseBuilder {
	b.expense.DueDate = dueDate
	return b
}

func (b *ExpenseBuilder) SetCreatedAt(createdAt time.Time) *ExpenseBuilder {
	b.expense.CreatedAt = createdAt
	return b
}

func (b *ExpenseBuilder) SetUpdatedAt(updatedAt time.Time) *ExpenseBuilder {
	b.expense.UpdatedAt = updatedAt
	return b
}

func (b *ExpenseBuilder) SetTypeID(typeID int) *ExpenseBuilder {
	b.expense.TypeID = typeID
	return b
}

func (b *ExpenseBuilder) Build() *Expense {
	return b.expense
}
