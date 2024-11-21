package middleware

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/config"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/OperatorSecurity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	log.Print("log authentication process")
	return func(c *gin.Context) {
		log.Print("check authorization")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}
		tokenString := parts[1]

		_, err := OperatorSecurity.Parse(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		isValid, err := config.ValidateToken(token)
		if err != nil || !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
