package main

import (
	"github.com/Ranzz02/auth-service/internal/db"
	"github.com/Ranzz02/auth-service/internal/models"
)

func main() {
	db := db.Init()

	db.AutoMigrate(&models.User{}, &models.Session{})
}
