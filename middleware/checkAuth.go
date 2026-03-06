package middleware

import (
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	TokenType string `json:"token_type"`
	UserId    string `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Next()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims := &UserClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		allowedIds := os.Getenv("DATA_MANAGER_ID")

		isValidToken := err == nil && token.Valid
		isAccessType := claims.TokenType == "access"
		isWhitelisted := strings.Contains(allowedIds, claims.UserId)

		if isValidToken && isAccessType && isWhitelisted {
			c.Set("user_id", claims.UserId)
		}

		c.Next()
	}
}
