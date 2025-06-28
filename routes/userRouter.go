package routes

import (
	controller "github.com/BhuwanSingh/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.GET("/users/signup", controller.signUp())
	incomingRoutes.GET("/users/login", controller.Login())
}
