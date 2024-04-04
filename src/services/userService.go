package services

import (
	"go-url-shortener/src/database"
	"go-url-shortener/src/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userService interface {
	CreateUser()
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc UserService) CreateUser(createUserInput *models.CreateUser) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error encrypting password")
	}
	createUserInput.Password = string(hash[:])
	user := &models.User{
		Email:    createUserInput.Email,
		Password: createUserInput.Password,
		Name:     createUserInput.Name,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return 500, nil
	}

	return 201, nil

}
