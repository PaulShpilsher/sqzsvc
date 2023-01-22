package main

import (
	"log"
	"sqzsvc/controllers"
	"sqzsvc/middlewares"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func registerRoutes(g *gin.Engine) {
	authController := &controllers.AuthController{}
	urlController := &controllers.UrlController{}

	g.GET("/:shortCode", urlController.GotoLongUrl)

	apiRoute := g.Group("/api")
	{
		route := apiRoute.Group("auth")
		route.POST("/register", authController.Register)
		route.POST("/login", authController.Login)
	}

	{
		route := apiRoute.Group("short-code")
		route.Use(middlewares.JwtAuthMiddleware())
		route.POST("/", urlController.CreateShortCode)
	}
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDb()

	r := gin.Default()
	registerRoutes(r)

	r.Run(":5555")
}
