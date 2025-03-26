package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpenseBuilder(t *testing.T) {
	dueDate := time.Now().Add(24 * time.Hour)
	tests := []struct {
		name        string
		id          string
		userID      string
		amount      float64
		description string
		category    string
		typeID      int
		paid        bool
		dueDate     time.Time
	}{
		{
			name:        "Valid expense",
			id:          "123",
			userID:      "user123",
			amount:      500.75,
			description: "Grocery shopping",
			category:    "Food",
			typeID:      2,
			paid:        true,
			dueDate:     dueDate,
		},
		{
			name:        "Unpaid expense",
			id:          "456",
			userID:      "user456",
			amount:      1200.00,
			description: "Rent",
			category:    "Housing",
			typeID:      2,
			paid:        false,
			dueDate:     dueDate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expense := NewExpenseBuilder().
				SetID(tt.id).
				SetUserID(tt.userID).
				SetAmount(tt.amount).
				SetDescription(tt.description).
				SetCategory(tt.category).
				SetTypeID(tt.typeID).
				SetPaid(tt.paid).
				SetDueDate(tt.dueDate).
				Build()

			assert.Equal(t, tt.id, expense.ID)
			assert.Equal(t, tt.userID, expense.UserID)
			assert.Equal(t, tt.amount, expense.Amount)
			assert.Equal(t, tt.description, expense.Description)
			assert.Equal(t, tt.category, expense.Category)
			assert.Equal(t, tt.typeID, expense.TypeID)
			assert.Equal(t, tt.paid, expense.Paid)
			assert.Equal(t, tt.dueDate, expense.DueDate)
			assert.False(t, expense.CreatedAt.IsZero())
			assert.False(t, expense.UpdatedAt.IsZero())
		})
	}
}

func TestExpenseBuilder_CustomDates(t *testing.T) {
	now := time.Now()
	dueDate := now.Add(7 * 24 * time.Hour)
	expense := NewExpenseBuilder().
		SetID("789").
		SetUserID("user789").
		SetAmount(350.25).
		SetDescription("Internet bill").
		SetCategory("Utilities").
		SetTypeID(2).
		SetPaid(false).
		SetDueDate(dueDate).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Build()

	assert.Equal(t, "789", expense.ID)
	assert.Equal(t, "user789", expense.UserID)
	assert.Equal(t, 350.25, expense.Amount)
	assert.Equal(t, "Internet bill", expense.Description)
	assert.Equal(t, "Utilities", expense.Category)
	assert.Equal(t, 2, expense.TypeID)
	assert.False(t, expense.Paid)
	assert.Equal(t, dueDate, expense.DueDate)
	assert.Equal(t, now, expense.CreatedAt)
	assert.Equal(t, now, expense.UpdatedAt)
}

func TestExpenseFactory(t *testing.T) {
	factory := NewExpenseFactory()
	transaction := factory.CreateTransaction()

	expense, ok := transaction.(*Expense)
	assert.True(t, ok, "Expected transaction to be of type *Expense")
	assert.NotNil(t, expense)
	assert.Equal(t, 2, expense.TypeID)
	assert.False(t, expense.CreatedAt.IsZero())
	assert.False(t, expense.UpdatedAt.IsZero())
}
