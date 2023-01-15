package main

import (
	"sqzsvc/controllers"

	"github.com/gin-gonic/gin"
)

func initUserRoutes(apiRoute *gin.RouterGroup) {

	userRoute := apiRoute.Group("user")

	userRoute.POST("/register", controllers.Register)

}

func main() {

	r := gin.Default()

	apiRoute := r.Group("/api")
	initUserRoutes(apiRoute)

	r.Run(":5555")
}
