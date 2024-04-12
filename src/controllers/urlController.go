package controllers

import (
	"errors"
	"go-url-shortener/src/common"
	"go-url-shortener/src/models"
	"go-url-shortener/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UrlController struct {
	logger  *zap.SugaredLogger
	service services.UrlService
}

func NewUrlController(l *zap.SugaredLogger, s services.UrlService) *UrlController {
	return &UrlController{l, s}
}

func (controller *UrlController) RedirectUrl(c *gin.Context) {
	hash := c.Param("hash")

	url, err := controller.service.GetUrl(hash)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

func (controller *UrlController) CreateUrl(c *gin.Context) {
	var body *models.CreateUrl
	c.Bind(&body)
	userId, exists := c.Get("sub")
	if !exists {
		controller.logger.Info("failed to find jwt's subject")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Cant find subject in jwt"})
		return
	}
	id, _ := strconv.Atoi(userId.(string))
	err := controller.service.CreateUrl(id, body)
	c.Status(http.StatusCreated)
	controller.logger.Infow("created new url", "userId", userId)
	switch {
	case err == nil:
		return
	case errors.As(err, &validator.ValidationErrors{}):
		controller.logger.Infof("failed validation")
		validationErrors := common.CollectErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}
