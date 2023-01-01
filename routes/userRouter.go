package routes

import (
	"github.com/tantm97/jwt-golang-project/middleware"

	controller "github.com/tantm97/jwt-golang-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	authorized := incomingRoutes.Group("/")
	authorized.Use(middleware.Authenticate())
	{
		authorized.GET("/users", controller.GetUsers())
		authorized.GET("/users/:user_id", controller.GetUser())
	}
	// incomingRoutes.Use(middleware.Authenticate())
	// incomingRoutes.GET("/users", controller.GetUsers())
	// incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
