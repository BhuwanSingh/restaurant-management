package routes

import (
	controller "github.com/BhuwanSingh/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Orders", controller.GetOrders())
	incomingRoutes.GET("/Orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/Orders", controller.CreateOrder())
	incomingRoutes.PATCH("/Orders/:order_id", controller.UpdateOrder())
}
