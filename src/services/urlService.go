package services

import (
	"errors"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/models"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UrlService struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUrlService(db *gorm.DB, l *zap.SugaredLogger) *UrlService {
	return &UrlService{db, l}
}

func (s *UrlService) CreateUrl(userId int, original_url string) error {
	result := s.db.Where("id = ?", userId).First(&models.User{})
	if result.RowsAffected == 0 {
		return errors.New("User don't exists")
	}
	hash, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 6)
	if err != nil {
		s.logger.Error("failed generating nanoid")
		return errors.New(apierrors.InternalServerError)
	}
	new_url := models.Url{
		OriginalUrl: original_url,
		Hash:        hash,
		UserId:      userId,
	}
	result = s.db.Create(&new_url)
	if result.Error != nil {
		s.logger.Errorf("failed to create new url", "user_id", userId)
		return errors.New(apierrors.InternalServerError)
	}
	s.logger.Infof("created new url", "user_id", userId)
	return nil

}
