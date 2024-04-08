package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) SetClaims(email string, exp_duration int) *MyCustomClaims {
	return &MyCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exp_duration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    email,
		},
	}
}

func (s *JwtService) GenerateToken(claims *MyCustomClaims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("Error while signing jwt")
	}
	return ss, nil
}
