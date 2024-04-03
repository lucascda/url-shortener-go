package controllers

import (
	"go-url-shortener/src/database"
	"go-url-shortener/src/models"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserController(c *gin.Context) {
	body := models.CreateUser{}

	c.Bind(&body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error encrypting password")
	}
	body.Password = string(hash[:])
	user := &models.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
