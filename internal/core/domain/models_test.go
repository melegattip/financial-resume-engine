package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// BaseModelBuilder es un builder para facilitar la creación de modelos base en pruebas
type BaseModelBuilder struct {
	model *BaseModel
}

func NewBaseModelBuilder() *BaseModelBuilder {
	return &BaseModelBuilder{
		model: &BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (b *BaseModelBuilder) WithID(id uint) *BaseModelBuilder {
	b.model.ID = id
	return b
}

func (b *BaseModelBuilder) WithCreatedAt(createdAt time.Time) *BaseModelBuilder {
	b.model.CreatedAt = createdAt
	return b
}

func (b *BaseModelBuilder) WithUpdatedAt(updatedAt time.Time) *BaseModelBuilder {
	b.model.UpdatedAt = updatedAt
	return b
}

func (b *BaseModelBuilder) Build() *BaseModel {
	return b.model
}

// UserBuilder es un builder para facilitar la creación de usuarios en pruebas
type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &User{
			BaseModel: BaseModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Active: true,
		},
	}
}

func (b *UserBuilder) WithID(id uint) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) WithFirstName(firstName string) *UserBuilder {
	b.user.FirstName = firstName
	return b
}

func (b *UserBuilder) WithLastName(lastName string) *UserBuilder {
	b.user.LastName = lastName
	return b
}

func (b *UserBuilder) WithActive(active bool) *UserBuilder {
	b.user.Active = active
	return b
}

func (b *UserBuilder) Build() *User {
	return b.user
}

// FinancialDataBuilder es un builder para facilitar la creación de datos financieros en pruebas
type FinancialDataBuilder struct {
	data *FinancialData
}

func NewFinancialDataBuilder() *FinancialDataBuilder {
	return &FinancialDataBuilder{
		data: &FinancialData{
			BaseModel: BaseModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}
}

func (b *FinancialDataBuilder) WithID(id uint) *FinancialDataBuilder {
	b.data.ID = id
	return b
}

func (b *FinancialDataBuilder) WithUserID(userID uint) *FinancialDataBuilder {
	b.data.UserID = userID
	return b
}

func (b *FinancialDataBuilder) WithAmount(amount float64) *FinancialDataBuilder {
	b.data.Amount = amount
	return b
}

func (b *FinancialDataBuilder) WithDescription(description string) *FinancialDataBuilder {
	b.data.Description = description
	return b
}

func (b *FinancialDataBuilder) WithCategory(category string) *FinancialDataBuilder {
	b.data.Category = category
	return b
}

func (b *FinancialDataBuilder) WithDate(date string) *FinancialDataBuilder {
	b.data.Date = date
	return b
}

func (b *FinancialDataBuilder) Build() *FinancialData {
	return b.data
}

func TestBaseModelBuilder(t *testing.T) {
	now := time.Now()
	model := NewBaseModelBuilder().
		WithID(1).
		WithCreatedAt(now).
		WithUpdatedAt(now).
		Build()

	assert.Equal(t, uint(1), model.ID)
	assert.Equal(t, now, model.CreatedAt)
	assert.Equal(t, now, model.UpdatedAt)
}

func TestUserBuilder(t *testing.T) {
	user := NewUserBuilder().
		WithID(1).
		WithEmail("test@example.com").
		WithPassword("password123").
		WithFirstName("John").
		WithLastName("Doe").
		WithActive(true).
		Build()

	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "password123", user.Password)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.True(t, user.Active)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestFinancialDataBuilder(t *testing.T) {
	data := NewFinancialDataBuilder().
		WithID(1).
		WithUserID(1).
		WithAmount(1000.50).
		WithDescription("Test transaction").
		WithCategory("Test category").
		WithDate("2024-03-25").
		Build()

	assert.Equal(t, uint(1), data.ID)
	assert.Equal(t, uint(1), data.UserID)
	assert.Equal(t, 1000.50, data.Amount)
	assert.Equal(t, "Test transaction", data.Description)
	assert.Equal(t, "Test category", data.Category)
	assert.Equal(t, "2024-03-25", data.Date)
	assert.False(t, data.CreatedAt.IsZero())
	assert.False(t, data.UpdatedAt.IsZero())
}
