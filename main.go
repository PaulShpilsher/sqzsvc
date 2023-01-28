package main

import (
	"log"
	authController "sqzsvc/controllers/auth"
	urlController "sqzsvc/controllers/url"
	"sqzsvc/middlewares"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func registerRoutes(g *gin.Engine) {
	g.GET("/:shortCode", urlController.RedirectShortCode)

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

	r.Run("localhost:5555")
}
