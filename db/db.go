package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file with godotenv: %w", err)
	}
	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Europe/Warsaw"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		log.Println("db is not initialized")
		err := initDB()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
