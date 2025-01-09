package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase() (*gorm.DB, error) {
	godotenv.Load(".env")

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSL := os.Getenv("DB_SSL")

	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=" + DB_SSL

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
