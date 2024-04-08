package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
	"go-url-shortener/src/validators.go"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

func InitServices(db *gorm.DB, v *validator.Validate) *controllers.UserController {

	userValidator := validators.NewUserValidator(v)
	userService := services.NewUserService(*userValidator, db)
	UserController := controllers.NewUserController(*userService)
	return UserController
}
