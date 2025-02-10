package service_interfaces

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type IJWTService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(tokenStr string) (*Claims, error)
}
