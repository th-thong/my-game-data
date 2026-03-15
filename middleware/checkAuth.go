package middleware

import (
	"context"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	AuthClient *auth.Client
}

var fbService *FirebaseService

func InitFirebase() error {
	cred := os.Getenv("FIREBASE_SERVICE_ACCOUNT")
	if cred == "" {
		return nil
	}
	opt := option.WithCredentialsJSON([]byte(cred))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return err
	}
	fbService = &FirebaseService{AuthClient: authClient}
	return nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if fbService == nil || fbService.AuthClient == nil {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenStr := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			c.Next()
			return
		}

		decodedToken, err := fbService.AuthClient.VerifyIDToken(c.Request.Context(), tokenStr)
		if err == nil {
			c.Set("user_id", decodedToken.UID)
		}

		c.Next()
	}
}
