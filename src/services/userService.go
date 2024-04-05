package services

import (
	"errors"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/validators.go"

	"go-url-shortener/src/database"
	"go-url-shortener/src/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userService interface {
	CreateUser()
}

type UserService struct {
	validator validators.UserValidator
}

func NewUserService(v validators.UserValidator) *UserService {
	return &UserService{v}
}

func (svc UserService) CreateUser(createUserInput *models.CreateUser) error {
	if err := svc.validator.ValidateCreateUser(createUserInput); err != nil {

		return err
	}

	var user models.User
	result := database.DB.Where("email = ?", createUserInput.Email).First(&user)
	if result.RowsAffected != 0 {
		return apierrors.UserAlreadyExistsError{}

	}
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error encrypting password")
	}
	createUserInput.Password = string(hash[:])
	user = models.User{
		Email:    createUserInput.Email,
		Password: createUserInput.Password,
		Name:     createUserInput.Name,
	}

	result = database.DB.Create(&user)
	if result.Error != nil {
		return errors.New("Error creating user in db")
	}

	return nil

}
