package routers

import (
	"rest_api/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrder)
	router.GET("/orders/:orderID", controllers.GetOrderbyID)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
