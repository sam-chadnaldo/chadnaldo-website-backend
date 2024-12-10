package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	UserCtx    = "user_id" // Вы можете хранить user ID в контексте для дальнейшего использования
)

// Initialize Firebase App (используйте один раз при старте приложения)
var firebaseAuthClient *auth.Client

func InitFirebase() error {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to initialize Firebase App: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return fmt.Errorf("failed to initialize Firebase Auth client: %v", err)
	}

	firebaseAuthClient = client
	return nil
}

// Middleware для проверки токена
func UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authHeader)
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Извлекаем токен из заголовка
		tokenString := strings.TrimPrefix(header, "Bearer ")
		if tokenString == header { // Если "Bearer " отсутствует
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Проверяем токен с помощью Firebase
		token, err := firebaseAuthClient.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("Failed to verify token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Добавляем данные пользователя в контекст
		c.Set(UserCtx, token.UID) // UID пользователя
		c.Set("user_email", token.Claims["email"]) // Email, если доступно
		c.Set("user_name", token.Claims["name"])   // Имя, если доступно
		c.Next()
	}
}
