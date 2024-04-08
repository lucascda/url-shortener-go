package services

import (
	"errors"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/validators.go"

	"gorm.io/gorm"

	"go-url-shortener/src/database"
	"go-url-shortener/src/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	validator validators.UserValidator
	db        *gorm.DB
}

func NewUserService(v validators.UserValidator, db *gorm.DB) *UserService {
	return &UserService{v, db}
}

func (svc *UserService) CreateUser(createUserInput *models.CreateUser) error {

	if err := svc.validator.ValidateCreateUser(createUserInput); err != nil {

		return err
	}

	var user models.User
	result := svc.db.Where("email = ?", createUserInput.Email).First(&user)
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
