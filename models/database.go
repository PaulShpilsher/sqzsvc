package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func InitDb() {
	dsn := os.Getenv("DB_CONNECTION")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to database: %s", err.Error())
		log.Fatal(err)
	}

	if err := database.AutoMigrate(&User{}, &UserUrl{}); err != nil {
		log.Fatalf("failed migrate database: %s", err.Error())
	}

	db = database
	log.Println("Connected to database")
}
