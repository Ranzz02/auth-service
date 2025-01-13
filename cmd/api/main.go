package main

import (
	"github.com/Ranzz02/auth-service/config"
	"github.com/Ranzz02/auth-service/internal/db"
	"github.com/Ranzz02/auth-service/internal/handlers"
	"github.com/Ranzz02/auth-service/internal/middleware"
	"github.com/Ranzz02/auth-service/internal/repositories"
	"github.com/Ranzz02/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
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

	// Instantiate goth
	goth.UseProviders(
		google.New(
			
		)
	)

	// Instantiate gin
	r := gin.New()

	//! Set middleware
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	db := db.Init()

	// Repositories
	authRepository := repositories.NewAuthRepository(db)

	// Services
	authService := services.NewAuthService(authRepository)

	baseName := r.Group("/")

	// Handlers
	handlers.NewAuthHandler(baseName, authRepository, authService) // Auth handler
	handlers.NewSessionHandler(baseName, authRepository)           // Session handler

	// Run API
	r.Run(envConfig.ServerHost + ":" + envConfig.ServerPort)
}
