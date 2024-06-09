package routers

import (
	"basic_trade/controllers"
	"basic_trade/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.AdminRegister) // url/auth/register
		authRouter.POST("/login", controllers.AdminLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetProduct)
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
	}

	return router
}