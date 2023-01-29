package models

import (
	"log"
	"sqzsvc/services/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	db = connectDb()
	log.Println("Connected to database")

	if err := db.AutoMigrate(
		&User{},
		&Url{},
		&Transition{}); err != nil {
		log.Fatalf("failed migrate database: %s", err.Error())
	}

	if err := InitSequenceGenerator(); err != nil {
		log.Fatalf("failed initialize number generator: %s", err.Error())
	}
}

func connectDb() *gorm.DB {
	database, err := gorm.Open(postgres.Open(config.DbDns), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to database: %s", err.Error())
	}
	return database
}
