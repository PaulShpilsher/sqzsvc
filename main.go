package main

import (
	"log"
	"sqzsvc/controllers/auth"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initUserRoutes(apiRoute *gin.RouterGroup) {

	userRoute := apiRoute.Group("user")

	userRoute.POST("/register", auth.Register)
	userRoute.POST("/login", auth.Login)
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDB()

	r := gin.Default()

	apiRoute := r.Group("/api")
	initUserRoutes(apiRoute)

	r.Run(":5555")
}
