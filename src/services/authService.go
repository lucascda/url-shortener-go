package services

import (
	"errors"
	"fmt"
	"go-url-shortener/src/models"
	"go-url-shortener/src/validators.go"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	jwtService JwtService
	validator  validators.AuthValidator
	db         *gorm.DB
}

func NewAuthService(db *gorm.DB, v validators.AuthValidator, jwtService JwtService) *AuthService {
	return &AuthService{jwtService, v, db}
}

func (s *AuthService) SignIn(signInInput *models.SignInInput) (string, error) {
	if err := s.validator.ValidateSignIn(signInInput); err != nil {
		return "", err
	}

	user := &models.User{}
	result := s.db.Where("email = ?", signInInput.Email).First(user)
	if result.RowsAffected == 0 {
		fmt.Print("PASSOU EMAIL")
		return "", errors.New("Unauthorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInInput.Password)); err != nil {
		return "", errors.New("Unauthorized")
	}

	claims := s.jwtService.SetClaims(signInInput.Email, 1)
	access_token, err := s.jwtService.GenerateToken(claims, []byte(os.Getenv("jwt_secret")))
	if err != nil {
		return "nil", err
	}

	return access_token, nil

}
