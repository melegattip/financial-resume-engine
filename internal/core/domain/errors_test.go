package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainErrors(t *testing.T) {
	// Verificar que los errores del dominio están definidos correctamente
	assert.NotNil(t, ErrInvalidTransactionType)
	assert.Equal(t, "invalid transaction type", ErrInvalidTransactionType.Error())
}

// DomainErrorBuilder es un builder para facilitar la creación de errores personalizados en pruebas
type DomainErrorBuilder struct {
	err error
}

func NewDomainErrorBuilder() *DomainErrorBuilder {
	return &DomainErrorBuilder{}
}

func (b *DomainErrorBuilder) WithMessage(message string) *DomainErrorBuilder {
	b.err = errors.New(message)
	return b
}

func (b *DomainErrorBuilder) WithError(err error) *DomainErrorBuilder {
	b.err = err
	return b
}

func (b *DomainErrorBuilder) Build() error {
	return b.err
}

func TestDomainErrorBuilder(t *testing.T) {
	// Probar la creación de un error personalizado
	customError := NewDomainErrorBuilder().
		WithMessage("test error").
		Build()

	assert.Error(t, customError)
	assert.Equal(t, "test error", customError.Error())

	// Probar la reutilización de un error existente
	existingError := ErrInvalidTransactionType
	wrappedError := NewDomainErrorBuilder().
		WithError(existingError).
		Build()

	assert.Error(t, wrappedError)
	assert.Equal(t, existingError.Error(), wrappedError.Error())
}
