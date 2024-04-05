package controllers

import (
	"go-url-shortener/src/common"
	apierrors "go-url-shortener/src/errors"
	"go-url-shortener/src/models"
	"go-url-shortener/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (userController *UserController) CreateUser(c *gin.Context) {
	body := models.CreateUser{}
	c.Bind(&body)
	err := userController.userService.CreateUser(&body)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, ok := err.(validator.ValidationErrors); ok {
			var validationErrors []common.ValidationError
			for _, err := range err.(validator.ValidationErrors) {
				validationError := common.NewValidationError(err.StructField(), err.Error(), err.Tag())

				validationErrors = append(validationErrors, *validationError)

			}
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
			return
		}

		if _, ok := err.(apierrors.UserAlreadyExistsError); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "Error creating user in db" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(201)
	return

}
