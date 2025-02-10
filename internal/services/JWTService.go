package services

import (
	"PPO_BMSTU/internal/services/service_interfaces"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

// JWT ключ из переменной окружения
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Структура для JWT-сервиса
type JWTService struct{}

// Создание нового экземпляра jwtService
func NewJWTService() service_interfaces.IJWTService {
	return &JWTService{}
}

// Генерация JWT токена для пользователя с UUID
func (s *JWTService) GenerateToken(userID uuid.UUID) (string, error) {
	// Создаем claims для JWT
	claims := service_interfaces.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // Токен на 2 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // Время выпуска токена
		},
	}

	// Создаем новый токен с указанными claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Проверка JWT
func (s *JWTService) ParseToken(tokenStr string) (*service_interfaces.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &service_interfaces.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*service_interfaces.Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
