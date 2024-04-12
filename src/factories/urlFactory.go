package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
	"go-url-shortener/src/validators"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitUrlController(l *zap.SugaredLogger, db *gorm.DB, v *validator.Validate) *controllers.UrlController {
	validator := validators.NewUrlValidator(v)
	service := services.NewUrlService(*validator, db, l)
	return controllers.NewUrlController(l, *service)
}
