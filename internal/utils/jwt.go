package utils

import (
	"time"

	"github.com/Ranzz02/auth-service/config"
	jwt "github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	Verify  TokenType = "verify"
)

func GenerateAccessToken(id string) (string, error) {
	lifespan := GetLifespan(Access)

	jti, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	config := config.NewEnvConfig()

	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["type"] = Access
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(lifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.TokenSecret)
}

func GenerateRefreshToken(id string) (string, string, error) {
	lifespan := GetLifespan(Refresh)

	jti, err := gonanoid.New()
	if err != nil {
		return "", "", err
	}

	config := config.NewEnvConfig()

	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["type"] = Refresh
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.TokenSecret)
	return jti, tokenString, err
}

func GetLifespan(Type TokenType) int {
	config := config.NewEnvConfig()

	switch Type {
	case Access:
		return config.TokenAccessTime
	case Refresh:
		return config.TokenRefreshTime
	case Verify:
		return config.TokenVerifyTime
	}
	return 0
}
