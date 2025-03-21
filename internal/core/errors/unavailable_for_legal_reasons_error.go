package errors

import "reflect"

type UnavailableForLegalReasons struct {
	message string
}

func NewUnavailableForLegalReasons(message string) UnavailableForLegalReasons {
	return UnavailableForLegalReasons{
		message: message,
	}
}

func (uflr UnavailableForLegalReasons) Error() string {
	return uflr.message
}

func IsUnavailableForLegalReasons(err error) bool {
	return reflect.TypeOf(err).String() == "errors.UnavailableForLegalReasons"
}

func IsUnavailableForLegalReasonsString(err string) bool {
	return err == "unavailable_for_legal_reasons_error"
}
