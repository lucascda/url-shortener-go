package validators

import (
	"go-url-shortener/src/models"

	"github.com/go-playground/validator/v10"
)

type UrlValidator struct {
	validator *validator.Validate
}

func NewUrlValidator(v *validator.Validate) *UrlValidator {
	return &UrlValidator{v}
}

func (v *UrlValidator) ValidateCreateUrl(createUrl *models.CreateUrl) error {
	return v.validator.Struct(createUrl)
}
