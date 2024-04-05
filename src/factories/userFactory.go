package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
	"go-url-shortener/src/validators.go"

	"github.com/go-playground/validator/v10"
)

var v = validator.New(validator.WithRequiredStructEnabled())
var userValidator = validators.NewUserValidator(v)
var userService = services.NewUserService(*userValidator)
var UserController = controllers.NewUserController(*userService)
