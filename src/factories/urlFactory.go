package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitUrlController(l *zap.SugaredLogger, db *gorm.DB) *controllers.UrlController {
	service := services.NewUrlService(db, l)
	return controllers.NewUrlController(l, *service)
}
