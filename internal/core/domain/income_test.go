package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIncomeBuilder(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		userID      string
		amount      float64
		description string
		category    string
		source      string
	}{
		{
			name:        "Valid income",
			id:          "123",
			userID:      "user123",
			amount:      1000.50,
			description: "Monthly salary",
			category:    "Salary",
			source:      "Employer XYZ",
		},
		{
			name:        "Zero amount income",
			id:          "456",
			userID:      "user456",
			amount:      0,
			description: "Bonus",
			category:    "Extra",
			source:      "Side project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			income := NewIncomeBuilder().
				SetID(tt.id).
				SetUserID(tt.userID).
				SetAmount(tt.amount).
				SetDescription(tt.description).
				SetCategory(tt.category).
				SetSource(tt.source).
				Build()

			assert.Equal(t, tt.id, income.ID)
			assert.Equal(t, tt.userID, income.UserID)
			assert.Equal(t, tt.amount, income.Amount)
			assert.Equal(t, tt.description, income.Description)
			assert.Equal(t, tt.category, income.Category)
			assert.Equal(t, tt.source, income.Source)
			assert.False(t, income.CreatedAt.IsZero())
			assert.False(t, income.UpdatedAt.IsZero())
		})
	}
}

func TestIncomeBuilder_CustomDates(t *testing.T) {
	now := time.Now()
	income := NewIncomeBuilder().
		SetID("789").
		SetUserID("user789").
		SetAmount(2500.75).
		SetDescription("Consulting fee").
		SetCategory("Freelance").
		SetSource("Client ABC").
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Build()

	assert.Equal(t, "789", income.ID)
	assert.Equal(t, "user789", income.UserID)
	assert.Equal(t, 2500.75, income.Amount)
	assert.Equal(t, "Consulting fee", income.Description)
	assert.Equal(t, "Freelance", income.Category)
	assert.Equal(t, "Client ABC", income.Source)
	assert.Equal(t, now, income.CreatedAt)
	assert.Equal(t, now, income.UpdatedAt)
}

func TestIncomeFactory(t *testing.T) {
	factory := NewIncomeFactory()
	transaction := factory.CreateTransaction()

	income, ok := transaction.(*Income)
	assert.True(t, ok, "Expected transaction to be of type *Income")
	assert.NotNil(t, income)
	assert.False(t, income.CreatedAt.IsZero())
	assert.False(t, income.UpdatedAt.IsZero())
}
