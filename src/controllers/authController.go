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
	if err := controller.service.SignIn(body); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
}
