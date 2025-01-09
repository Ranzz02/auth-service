package main

import (
	"github.com/Ranzz02/auth-service/config"
	"github.com/Ranzz02/auth-service/internal/db"
	"github.com/Ranzz02/auth-service/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	envConfig := config.NewEnvConfig()
	logLevel, err := logrus.ParseLevel(envConfig.LogLevel)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	if envConfig.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.Use(gin.Recovery())

	db := db.Init()

	// Repositories
	authRepository := repositories.NewAuthRepository(db)

}
