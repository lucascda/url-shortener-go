package controllers

import (
	"errors"
	"go-url-shortener/src/common"
	"go-url-shortener/src/database"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/models"
	"go-url-shortener/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserController struct {
	logger      *zap.SugaredLogger
	userService services.UserService
}

func NewUserController(l *zap.SugaredLogger, service services.UserService) *UserController {
	return &UserController{l, service}
}

func (controller *UserController) Profile(c *gin.Context) {
	issuer, exists := c.Get("issuer")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Can't find issuer in request"})
		return
	}
	user := &models.User{}
	result := database.DB.Where("email = ?", issuer).First(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Cant find user"})
		return
	}
	c.JSON(200, gin.H{"me": user.ID})
	return
}

func (controller *UserController) CreateUser(c *gin.Context) {
	body := models.CreateUser{}
	c.Bind(&body)
	err := controller.userService.CreateUser(&body)
	handleError(c, controller.logger, err)

	c.Status(201)
	return

}

func handleError(c *gin.Context, logger *zap.SugaredLogger, err error) {
	switch {
	case err == nil:
		return
	case errors.Is(err, &validator.InvalidValidationError{}):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	case errors.As(err, &validator.ValidationErrors{}):
		logger.Infof("failed validation")
		var validationErrors []common.ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			validationError := common.NewValidationError(err.StructField(), err.Error(), err.Tag())

			validationErrors = append(validationErrors, *validationError)

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
		return
	case errors.Is(err, apierrors.UserAlreadyExistsError{}):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
