package main

import (
	authController "sqzsvc/controllers/auth"
	urlController "sqzsvc/controllers/url"
	"sqzsvc/middlewares"
	"sqzsvc/models"
	"sqzsvc/services/config"

	"github.com/gin-gonic/gin"
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
	config.InitConfig()

	models.InitDb()

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	registerRoutes(r)
	r.Run(":5555")
}
