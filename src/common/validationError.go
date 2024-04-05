package common

type ValidationError struct {
	Field string
	Msg   string
	Tag   string
}

func NewValidationError(field string, msg string, tag string) *ValidationError {
	return &ValidationError{field, msg, tag}
}
