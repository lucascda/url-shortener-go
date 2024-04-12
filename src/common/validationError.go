package common

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	Field string
	Msg   string
	Tag   string
}

func NewValidationError(field string, msg string, tag string) *ValidationError {
	return &ValidationError{field, msg, tag}
}

func CollectErrors(err error) []ValidationError {
	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		validationError := NewValidationError(err.StructField(), err.Error(), err.Tag())
		validationErrors = append(validationErrors, *validationError)
	}

	return validationErrors
}
