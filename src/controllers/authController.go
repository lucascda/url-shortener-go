package controllers

import (
	"go-url-shortener/src/models"
	"go-url-shortener/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(s services.AuthService) *AuthController {
	return &AuthController{s}
}

func (controller *AuthController) SignIn(c *gin.Context) {
	body := &models.SignInInput{}
	c.Bind(body)
	access_token, err := controller.service.SignIn(body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"access_token": access_token})
	return
}
