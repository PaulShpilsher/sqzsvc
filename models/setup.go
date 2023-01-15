package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() {
	Database = NewDB()
}

func NewDB() *gorm.DB {
	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	db.AutoMigrate(&User{})
	return db
}
