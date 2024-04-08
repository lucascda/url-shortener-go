package services

import (
	"errors"
	"fmt"
	"go-url-shortener/src/models"
	"go-url-shortener/src/validators.go"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

type AuthService struct {
	validator validators.AuthValidator
	db        *gorm.DB
}

func NewAuthService(db *gorm.DB, v validators.AuthValidator) *AuthService {
	return &AuthService{v, db}
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

	claims := MyCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("random_string"))
	if err != nil {

		return "", errors.New("Erro ao assinar jwt")
	}

	return ss, nil

}
