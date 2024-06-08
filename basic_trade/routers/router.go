package routers

import (
	"basic_trade/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.AdminRegister) // url/auth/register
		authRouter.POST("/login", controllers.AdminLogin)
	}

	// router.POST("/products", controllers.CreateProduct)

	return router
}
