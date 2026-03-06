package middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
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

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		pubKeyB64 := os.Getenv("JWT_PUBLIC_KEY_B64")
		pubKeyPEM, err := base64.StdEncoding.DecodeString(pubKeyB64)
		if err != nil {
			c.Next()
			return
		}

		pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyPEM)
		if err != nil {
			c.Next()
			return
		}

		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return pubKey, nil
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