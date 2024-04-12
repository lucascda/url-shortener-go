package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
	"go-url-shortener/src/validators"

	"go.uber.org/zap"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

func InitServices(l *zap.SugaredLogger, db *gorm.DB, v *validator.Validate) *controllers.UserController {

	userValidator := validators.NewUserValidator(v)
	userService := services.NewUserService(l, *userValidator, db)
	UserController := controllers.NewUserController(l, *userService)
	return UserController
}
