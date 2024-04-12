package common

import (
	"go-url-shortener/src/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(header, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong authorization token format"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		claims, valid := services.NewJwtService().ParseToken(token)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		Logger.Info(claims)
		sub, err := claims.GetSubject()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing subject"})
			c.Abort()
			return
		}
		c.Set("sub", sub)
		c.Next()
	}
}
