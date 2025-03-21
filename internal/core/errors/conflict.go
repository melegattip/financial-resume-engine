package errors

type Conflict struct {
	message string
}

func NewConflict(message string) Conflict {
	return Conflict{
		message: message,
	}
}

func (e Conflict) Error() string {
	return e.message
}
