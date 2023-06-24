package model

type Error struct {
	Message string
	Status  int
}

func (e Error) Error() string {
	return e.Message
}

func NewError(status int, message string) Error {
	return Error{
		Status:  status,
		Message: message,
	}
}
