package main

import (
	"github.com/Ranzz02/auth-service/config"
	"github.com/Ranzz02/auth-service/internal/db"
	"github.com/Ranzz02/auth-service/internal/handlers"
	"github.com/Ranzz02/auth-service/internal/middleware"
	"github.com/Ranzz02/auth-service/internal/repositories"
	"github.com/Ranzz02/auth-service/internal/services"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Configure everything
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

	// Instantiate gin
	r := gin.New()

	//! Set middleware
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	// Instantiate mail
	utils.NewEmailHandler()

	// Init database connection
	db := db.Init()

	// Repositories
	authRepository := repositories.NewAuthRepository(db)
	userRepository := repositories.NewUserRepository(db)

	// Services
	authService := services.NewAuthService(authRepository)

	// Router Groups
	baseRouter := r.Group("/", middleware.AuthMiddleware)                 // "/"" group (base)
	authRouter := r.Group("/auth")                                        // "/auth" group
	sessionRouter := r.Group("/auth/sessions", middleware.AuthMiddleware) // "/auth/sessions" group

	// Handlers
	handlers.NewAuthHandler(authRouter, authRepository, authService) // Auth handler
	handlers.NewSessionHandler(sessionRouter, authRepository)        // Session handler
	handlers.NewUserHandler(baseRouter, userRepository)              // User handler

	// Run API
	r.Run(envConfig.ServerHost + ":" + envConfig.ServerPort)
}
