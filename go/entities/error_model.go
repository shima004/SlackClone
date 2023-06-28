package entities

type errInvalidPassword struct {
}

func (e *errInvalidPassword) Error() string {
	return "invalid password"
}

var ErrInvalidPassword = &errInvalidPassword{}

type errDataNotFound struct {
}

func (e *errDataNotFound) Error() string {
	return "data not found"
}

var ErrDataNotFound = &errDataNotFound{}

type errChannelNotFound struct {
}

func (e *errChannelNotFound) Error() string {
	return "channel not found"
}

var ErrChannelNotFound = &errChannelNotFound{}

type errValidation struct {
	message string
}

func (e *errValidation) Error() string {
	return e.message
}

func NewValidationError(message string) error {
	return &errValidation{message: message}
}
