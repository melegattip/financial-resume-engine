// Package domain defines the core business entities and their behavior
package domain

import "time"

// Expense represents a financial expense transaction
type Expense struct {
	BaseTransaction
	TypeID  int       `json:"type_id"`
	Paid    bool      `json:"paid,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
}

// ExpenseFactory implements TransactionFactory for creating expense transactions
type ExpenseFactory struct{}

// NewExpenseFactory creates a new ExpenseFactory instance
func NewExpenseFactory() *ExpenseFactory {
	return &ExpenseFactory{}
}

// CreateTransaction creates a new expense transaction with default values
func (f *ExpenseFactory) CreateTransaction() Transaction {
	return &Expense{
		BaseTransaction: BaseTransaction{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		TypeID: 2, // 2 represents "expense"
	}
}

// ExpenseBuilder implements the builder pattern for creating Expense instances
type ExpenseBuilder struct {
	expense *Expense
}

// NewExpenseBuilder creates a new ExpenseBuilder instance with default values
func NewExpenseBuilder() *ExpenseBuilder {
	return &ExpenseBuilder{
		expense: &Expense{
			BaseTransaction: BaseTransaction{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			TypeID: 2, // 2 represents "expense"
		},
	}
}

// SetID sets the expense ID and returns the builder
func (b *ExpenseBuilder) SetID(id string) *ExpenseBuilder {
	b.expense.ID = id
	return b
}

// SetUserID sets the user ID and returns the builder
func (b *ExpenseBuilder) SetUserID(userID string) *ExpenseBuilder {
	b.expense.UserID = userID
	return b
}

// SetAmount sets the expense amount and returns the builder
func (b *ExpenseBuilder) SetAmount(amount float64) *ExpenseBuilder {
	b.expense.Amount = amount
	return b
}

// SetDescription sets the expense description and returns the builder
func (b *ExpenseBuilder) SetDescription(description string) *ExpenseBuilder {
	b.expense.Description = description
	return b
}

// SetCategory sets the expense category and returns the builder
func (b *ExpenseBuilder) SetCategory(category string) *ExpenseBuilder {
	b.expense.Category = category
	return b
}

// SetPaid sets the paid status and returns the builder
func (b *ExpenseBuilder) SetPaid(paid bool) *ExpenseBuilder {
	b.expense.Paid = paid
	return b
}

// SetDueDate sets the due date and returns the builder
func (b *ExpenseBuilder) SetDueDate(dueDate time.Time) *ExpenseBuilder {
	b.expense.DueDate = dueDate
	return b
}

// SetCreatedAt sets the creation timestamp and returns the builder
func (b *ExpenseBuilder) SetCreatedAt(createdAt time.Time) *ExpenseBuilder {
	b.expense.CreatedAt = createdAt
	return b
}

// SetUpdatedAt sets the last update timestamp and returns the builder
func (b *ExpenseBuilder) SetUpdatedAt(updatedAt time.Time) *ExpenseBuilder {
	b.expense.UpdatedAt = updatedAt
	return b
}

// SetTypeID sets the transaction type ID and returns the builder
func (b *ExpenseBuilder) SetTypeID(typeID int) *ExpenseBuilder {
	b.expense.TypeID = typeID
	return b
}

// Build creates and returns a new Expense instance
func (b *ExpenseBuilder) Build() *Expense {
	return b.expense
}
