package services

import (
	"errors"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/validators.go"

	"go.uber.org/zap"

	"gorm.io/gorm"

	"go-url-shortener/src/database"
	"go-url-shortener/src/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	logger    *zap.SugaredLogger
	validator validators.UserValidator
	db        *gorm.DB
}

func NewUserService(l *zap.SugaredLogger, v validators.UserValidator, db *gorm.DB) *UserService {
	return &UserService{l, v, db}
}

func (s *UserService) CreateUser(createUserInput *models.CreateUser) error {

	if err := s.validator.ValidateCreateUser(createUserInput); err != nil {

		return err
	}

	var user models.User
	result := s.db.Where("email = ?", createUserInput.Email).First(&user)
	if result.RowsAffected != 0 {
		s.logger.Infow("user already exists", "email", createUserInput.Email)
		return apierrors.UserAlreadyExistsError{}

	}
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserInput.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Infow("failed to encrypt password", "email", createUserInput.Email)
		return errors.New("failed encrypting password")
	}
	createUserInput.Password = string(hash[:])
	user = models.User{
		Email:    createUserInput.Email,
		Password: createUserInput.Password,
		Name:     createUserInput.Name,
	}

	result = database.DB.Create(&user)
	if result.Error != nil {
		s.logger.Errorf("failed to save user in database", "email", createUserInput.Email)
		return errors.New("Error creating user in db")
	}

	return nil

}
