package main

import (
	"log"
	"sqzsvc/controllers"
	"sqzsvc/middlewares"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDb()

	r := gin.Default()
	{
		apiRoute := r.Group("/api")
		{
			authController := &controllers.AuthController{}
			route := apiRoute.Group("auth")
			route.POST("/register", authController.Register)
			route.POST("/login", authController.Login)
		}

		{
			userController := &controllers.UserController{}
			route := apiRoute.Group("user")
			route.Use(middlewares.JwtAuthMiddleware())
			route.GET("/current", userController.CurrentUser)
		}
	}

	r.Run(":5555")
}
