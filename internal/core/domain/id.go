package domain

import (
	"github.com/google/uuid"
)

// NewID genera un nuevo ID único usando UUID v4
func NewID() string {
	return uuid.New().String()
}
