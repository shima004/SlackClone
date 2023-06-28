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

type errUnauthorized struct {
}

func (e *errUnauthorized) Error() string {
	return "unauthorized"
}

var ErrUnauthorized = &errUnauthorized{}

type errValidation struct {
	message string
}

var ErrValidation = &errValidation{}

func (e *errValidation) Error() string {
	return e.message
}

func NewValidationError(message string) error {
	return &errValidation{message: message}
}
