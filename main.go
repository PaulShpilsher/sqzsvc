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
			route := apiRoute.Group("auth")
			authController := &controllers.AuthController{}
			route.POST("/register", authController.Register)
			route.POST("/login", authController.Login)
		}

		{
			route := apiRoute.Group("user")
			userController := &controllers.UserController{}
			route.Use(middlewares.JwtAuthMiddleware())
			route.GET("/current", userController.CurrentUser)
			route.POST("/url", userController.RegisterLongUrl)
		}
	}

	r.Run(":5555")
}
