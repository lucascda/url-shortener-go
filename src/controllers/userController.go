package controllers

import (
	"go-url-shortener/src/models"
	"go-url-shortener/src/services"

	"github.com/gin-gonic/gin"
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
	result, err := userController.userService.CreateUser(&body)
	if err != nil {
		c.Status(400)
		return
	}
	c.Status(result)

}
