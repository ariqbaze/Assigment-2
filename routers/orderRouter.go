package routers

import (
	"Assigment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/order", controllers.GetAllOrder)

	router.POST("/order", controllers.CreateOrder)

	router.PUT("/order/:orderId", controllers.UpdateOrder)

	router.DELETE("/order/:orderId", controllers.DeleteOrder)
	return router
}
