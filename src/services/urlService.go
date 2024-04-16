package services

import (
	"errors"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/models"
	"go-url-shortener/src/validators"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UrlService struct {
	validator validators.UrlValidator
	db        *gorm.DB
	logger    *zap.SugaredLogger
}

func NewUrlService(v validators.UrlValidator, db *gorm.DB, l *zap.SugaredLogger) *UrlService {
	return &UrlService{v, db, l}
}

func (s *UrlService) GetUrl(hash string) (string, error) {

	url := models.Url{}
	result := s.db.Where("hash = ?", hash).First(&url)
	s.logger.Infow("url found", "url", url)
	if result.RowsAffected == 0 {
		s.logger.Info("url not found")
		return "", errors.New("Url not found")
	}

	return url.OriginalUrl, nil
}

func (s *UrlService) ListUrls(userId int) (any, error) {

	result := s.db.Where("id = ?", userId).First(&models.User{})
	if result.RowsAffected == 0 {
		return nil, errors.New("User don't exists")
	}
	urls := []models.Url{}
	result = s.db.Find(&urls)
	listAll := []*models.ListAllUrls{}

	for _, url := range urls {
		listAll = append(listAll, &models.ListAllUrls{
			ID:           url.ID,
			Original_url: url.OriginalUrl,
			Hash:         url.Hash,
			CreatedAt:    url.CreatedAt,
			UpdatedAt:    url.UpdatedAt,
		})
	}

	return listAll, nil
}

func (s *UrlService) CreateUrl(userId int, createUrl *models.CreateUrl) error {
	if err := s.validator.ValidateCreateUrl(createUrl); err != nil {
		return err
	}

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
		OriginalUrl: createUrl.Original_url,
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
