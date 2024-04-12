package validators

import (
	"go-url-shortener/src/models"

	"github.com/go-playground/validator/v10"
)

type AuthValidator struct {
	validate *validator.Validate
}

func NewAuthValidator(v *validator.Validate) *AuthValidator {
	return &AuthValidator{v}
}

func (v *AuthValidator) ValidateSignIn(signInInput *models.SignInInput) error {
	return v.validate.Struct(signInInput)
}
