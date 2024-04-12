package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
	"go-url-shortener/src/validators"

	"go.uber.org/zap"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

func InitAuthController(l *zap.SugaredLogger, db *gorm.DB, v *validator.Validate) *controllers.AuthController {
	authValidator := validators.NewAuthValidator(v)
	jwtService := services.NewJwtService()
	authService := services.NewAuthService(l, db, *authValidator, *jwtService)
	authController := controllers.NewAuthController(*authService)
	return authController
}
