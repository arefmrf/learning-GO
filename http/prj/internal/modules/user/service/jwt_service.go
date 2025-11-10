package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("super_secret_key")

type JWTService interface {
	GenerateToken(username string) (string, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(username string) (string, error) {
	claims := CustomClaims{
		Username: fmt.Sprint(username),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
