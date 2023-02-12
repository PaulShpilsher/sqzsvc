package main

import (
	"fmt"
	"log"
	authController "sqzsvc/controllers/auth"
	urlController "sqzsvc/controllers/url"
	"sqzsvc/models"
	"sqzsvc/services/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
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
		// route.Use(middlewares.JwtAuthMiddleware())
		route.POST("/", urlController.CreateShortCode)
	}

	// serve static files
	g.Use(static.Serve("/", static.LocalFile("./www", true)))
	g.NoRoute(func(c *gin.Context) { // fallback
		c.File("./www/index.html")
	})
}

func main() {
	config.InitConfig()

	models.InitDb()

	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true //IMPORTANT: NOT FOR PROD
	router.Use(cors.New(corsConfig))
	// router.Use(cors.Default())

	registerRoutes(router)

	serverAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Server started at http://%s ...\n", serverAddress)
	router.Run(serverAddress)

}
