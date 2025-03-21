package errors_test

import (
	"testing"

	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewUnavailableForLegalReasons(t *testing.T) {
	message := "Unavailable for legal reasons error"
	err := errors.NewUnavailableForLegalReasons(message)

	assert.Equal(t, message, err.Error())
}

func TestIsUnavailableForLegalReasons(t *testing.T) {
	err := errors.NewUnavailableForLegalReasons("error")
	assert.True(t, errors.IsUnavailableForLegalReasons(err))
}

func TestIsUnavailableForLegalReasonsString(t *testing.T) {
	assert.True(t, errors.IsUnavailableForLegalReasonsString("unavailable_for_legal_reasons_error"))
}
