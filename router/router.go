package router

import (
	"assignment2/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	orderRouter := gin.Default()

	{
		orderRouter.POST("/orders", controller.CreateOrders)
		orderRouter.GET("/orders", controller.GetOrders)
		orderRouter.PUT("/orders/:orderId", controller.UpdateOrders)
		//orderRouter.PUT("/orders/:orderId",controller.CreateOrders)
		orderRouter.DELETE("/orders/:orderId", controller.DeleteOrder)
	}

	return orderRouter
}
