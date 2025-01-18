package utils

import (
	"fmt"
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
	return token.SignedString([]byte(config.TokenSecret))
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
	tokenString, err := token.SignedString([]byte(config.TokenSecret))
	return jti, tokenString, err
}

// Generate Verify Token
//
// Generates a token that is sent to the email of a user on creation of the user and when user requests a new verification email.
func GenerateVerifyToken(id string, code string) (string, error) {
	lifespan := GetLifespan(Verify)

	jti, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	config := config.NewEnvConfig()

	claims := jwt.MapClaims{}
	claims["jti"] = jti
	claims["type"] = Verify
	claims["sub"] = id
	claims["code"] = code
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.TokenSecret))
	return tokenString, err
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

func ExtractClaims(token string) jwt.MapClaims {
	config := config.NewEnvConfig()
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		return nil
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil
	}

	return claims
}

func VerifyToken(token string, tokenType TokenType) bool {
	config := config.NewEnvConfig()
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		return false
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return false
	}

	if claims["type"] == string(tokenType) {
		return false
	}

	return true
}
