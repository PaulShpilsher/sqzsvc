package main

import (
	"log"
	"sqzsvc/controllers/auth"
	"sqzsvc/middlewares"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initAuthRoutes(apiRoute *gin.RouterGroup) {
	route := apiRoute.Group("auth")
	route.POST("/register", auth.Register)
	route.POST("/login", auth.Login)
}

func initAdminRoutes(apiRoute *gin.RouterGroup) {
	route := apiRoute.Group("user")
	route.Use(middlewares.JwtAuthMiddleware())
	route.GET("/current", auth.CurrentUser)
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDB()

	r := gin.Default()

	apiRoute := r.Group("/api")
	initAuthRoutes(apiRoute)
	initAdminRoutes(apiRoute)

	r.Run(":5555")
}
