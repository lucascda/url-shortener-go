package main

import (
	"go-url-shortener/src/database"
	"go-url-shortener/src/factories"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDB()
}

func main() {

	userController := factories.InitServices(database.DB)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/users", userController.CreateUser)

	r.Run()
}
