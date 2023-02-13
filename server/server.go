package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	authController "sqzsvc/controllers/auth"
	urlController "sqzsvc/controllers/url"
	"sqzsvc/models"
	"sqzsvc/services/config"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const readTimeout = 10 * time.Second
const writeTimeout = 10 * time.Second

func Serve() {
	config.InitConfig()

	models.InitDb()

	router := createRouter()
	registerRoutes(router)
	startServer(router)
}

func startServer(router *gin.Engine) {
	var serverAddress string
	if config.Docker {
		log.Println("Running in docker")
		serverAddress = fmt.Sprintf("0.0.0.0:%d", config.Port)
	} else {
		serverAddress = fmt.Sprintf("%s:%d", config.Host, config.Port)
	}

	server := &http.Server{
		Addr:         serverAddress,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("Starting server at http://%s\n", serverAddress)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func createRouter() *gin.Engine {
	// init router
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New() // gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true //IMPORTANT: NOT FOR PROD
	router.Use(cors.New(corsConfig))  // router.Use(cors.Default())

	return router
}

func registerRoutes(router *gin.Engine) {
	// redirect
	router.GET("/:shortCode", urlController.RedirectShortCode)

	apiRoute := router.Group("/api")
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
	router.Use(static.Serve("/", static.LocalFile("./www", true)))
	router.NoRoute(func(c *gin.Context) { // fallback
		c.File("./www/index.html")
	})
}
