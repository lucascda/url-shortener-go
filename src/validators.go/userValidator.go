package validators

import (
	"fmt"
	"go-url-shortener/src/models"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator(validate *validator.Validate) *UserValidator {
	return &UserValidator{validate: validate}
}

func (uv *UserValidator) ValidateCreateUser(user *models.CreateUser) error {
	fmt.Print("call ValidateCreateUser")
	return uv.validate.Struct(user)
}
