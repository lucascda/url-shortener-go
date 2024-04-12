package services

import (
	"errors"
	"go-url-shortener/src/models"
	"go-url-shortener/src/validators"
	"os"
	"strconv"

	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	logger     *zap.SugaredLogger
	jwtService JwtService
	validator  validators.AuthValidator
	db         *gorm.DB
}

func NewAuthService(l *zap.SugaredLogger, db *gorm.DB, v validators.AuthValidator, jwtService JwtService) *AuthService {
	return &AuthService{l, jwtService, v, db}
}

func (s *AuthService) SignIn(signInInput *models.SignInInput) (string, error) {
	if err := s.validator.ValidateSignIn(signInInput); err != nil {
		return "", err
	}

	user := &models.User{}
	result := s.db.Where("email = ?", signInInput.Email).First(user)
	if result.RowsAffected == 0 {
		s.logger.Infow("failed to find user", "email", signInInput.Email)
		return "", errors.New("Unauthorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInInput.Password)); err != nil {
		s.logger.Infow("failed comparing passwords", "user", user.ID, "email", user.Email)
		return "", errors.New("Unauthorized")
	}

	claims := s.jwtService.SetClaims("go-url-api", strconv.Itoa(int(user.ID)), 1)

	access_token, err := s.jwtService.GenerateToken(claims, []byte(os.Getenv("jwt_secret")))
	if err != nil {
		s.logger.Infow("failed generating token", "user", user.ID, "email", user.Email)
		return "", err
	}

	return access_token, nil

}
