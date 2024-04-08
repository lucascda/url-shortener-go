package services

import (
	"errors"
	"go-url-shortener/src/models"
	"go-url-shortener/src/validators.go"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type AuthService struct {
	validator validators.AuthValidator
	db        *gorm.DB
}

func NewAuthService(db *gorm.DB, v validators.AuthValidator) *AuthService {
	return &AuthService{v, db}
}

func (s *AuthService) SignIn(signInInput *models.SignInInput) error {
	if err := s.validator.ValidateSignIn(signInInput); err != nil {
		return err
	}

	user := &models.User{}
	result := s.db.Where("email = ?", signInInput.Email).First(user)
	if result.RowsAffected == 0 {
		return errors.New("Unauthorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInInput.Password)); err != nil {
		return errors.New("Unauthorized")
	}

	return nil

}
