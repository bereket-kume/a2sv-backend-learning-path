package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(useranme, role string) (string, error)
}

type JWTServiceImpl struct {
	SecretKey string
}

func NewJWTService(secretKey string) *JWTServiceImpl {
	return &JWTServiceImpl{SecretKey: secretKey}
}

func (j *JWTServiceImpl) GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(j.SecretKey))
}
