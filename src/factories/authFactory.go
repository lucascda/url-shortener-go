package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
	"go-url-shortener/src/validators.go"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

func InitAuthController(db *gorm.DB, v *validator.Validate) *controllers.AuthController {
	authValidator := validators.NewAuthValidator(v)
	jwtService := services.NewJwtService()
	authService := services.NewAuthService(db, *authValidator, *jwtService)
	authController := controllers.NewAuthController(*authService)
	return authController
}
