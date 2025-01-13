package db

import (
	"fmt"
	"log"

	"github.com/Ranzz02/auth-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase() (*gorm.DB, error) {
	config := config.NewEnvConfig()

	dsn := "host=" + config.DBHost + " user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " port=" + config.DBPort + " sslmode=" + config.DBSSLMode

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatalln("connection error: ", err)
		return nil, err
	}
	fmt.Println("We are connected to the database ")

	return db, nil
}

func Init() *gorm.DB {
	db, err := ConnectDatabase()

	if err != nil {
		return nil
	}
	return db
}
