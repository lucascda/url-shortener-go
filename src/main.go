package main

import (
	"go-url-shortener/src/common"
	"go-url-shortener/src/database"
	"go-url-shortener/src/factories"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	common.Logger = common.InitLogger()
	common.Logger.Info("Loaded logger")
	prometheus.MustRegister(common.RequestsTotal)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDB()
}

func main() {

	v := validator.New(validator.WithRequiredStructEnabled())
	userController := factories.InitServices(common.Logger, database.DB, v)
	authController := factories.InitAuthController(common.Logger, database.DB, v)
	urlController := factories.InitUrlController(common.Logger, database.DB)
	r := gin.Default()
	r.Use(common.PrometheusMiddleware())

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/users", userController.CreateUser)
	r.POST("/signin", authController.SignIn)

	protected := r.Group("/protected")
	protected.Use(common.JwtAuthMiddleware())
	{
		protected.GET("", userController.Profile)
	}
	urls := r.Group("/urls")
	urls.Use(common.JwtAuthMiddleware())
	{
		urls.POST("", urlController.CreateUrl)
	}
	r.Run()
}
