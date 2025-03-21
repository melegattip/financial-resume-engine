package errors

import "reflect"

type SkippableError struct {
	message string
}

func NewSkippableError(message string) SkippableError {
	return SkippableError{
		message: message,
	}
}

func (e SkippableError) Error() string {
	return e.message
}

func IsSkippableError(err error) bool {
	return reflect.TypeOf(err).String() == "errors.SkippableError"
}
