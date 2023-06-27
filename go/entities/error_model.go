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
