package errors_test

import (
	"testing"

	"github.com/melegattip/financial-resume-engine/internal/core/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewSkippableError(t *testing.T) {
	message := "This is an skippable error"
	err := errors.NewSkippableError(message)
	assert.Equal(t, message, err.Error())
}

func TestIsSkippableError(t *testing.T) {
	err := errors.NewSkippableError("This is a validation error")

	assert.True(t, errors.IsSkippableError(err))
}
