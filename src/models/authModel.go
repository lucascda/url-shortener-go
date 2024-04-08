package models

type SignInInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
