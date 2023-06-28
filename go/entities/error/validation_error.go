package cerror

type ValidationError struct {
	Errors map[string][]string
}

func (e *ValidationError) Error() string {
	message := ""
	for paramName, reasons := range e.Errors {
		for _, reason := range reasons {
			message += paramName + " " + reason + "\n"
		}
	}
	return message
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		Errors: make(map[string][]string),
	}
}

func (e *ValidationError) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ValidationError) GetErrors() map[string][]string {
	return e.Errors
}

func (e *ValidationError) Add(paramName string, reason string) {
	if e.Errors[paramName] == nil {
		e.Errors[paramName] = []string{}
	}
	e.Errors[paramName] = append(e.Errors[paramName], reason)
}

func (e *ValidationError) Is(target error) bool {
	_, ok := target.(*ValidationError)
	return ok
}
