package controllersApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Логин — генерируем токен
func (s *ServicesAPI) Login(c *gin.Context) {
	var loginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	// Парсим запрос
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Логика проверки логина и пароля, например через сервис
	judge, token, err := s.Services.JudgeService.Login(loginRequest.Login, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем токен
	c.JSON(http.StatusOK, gin.H{
		"judge": judge,
		"token": token,
	})
}

// Логаут — фактически, просто отсутствие токена на клиенте
func (s *ServicesAPI) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

// JWTMiddleware проверяет, существует ли токен и является ли он валидным
func (s *ServicesAPI) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Извлекаем токен из заголовка
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Malformed token"})
			c.Abort()
			return
		}

		// Токен идет в формате "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Парсим токен
		tokenStr := tokenParts[1]
		claims, err := s.Services.JWTService.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Добавляем данные пользователя в контекст запроса
		c.Set("user_id", claims.UserID)

		// Переходим к следующему обработчику
		c.Next()
	}
}
