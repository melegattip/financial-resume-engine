package errors

import "reflect"

type BadRequest struct {
	message string
}

func NewBadRequest(message string) *BadRequest {
	return &BadRequest{
		message: message,
	}
}

func (e *BadRequest) Error() string {
	return e.message
}

func IsBadRequestError(err error) bool {
	return reflect.TypeOf(err).String() == "*errors.BadRequest"
}
