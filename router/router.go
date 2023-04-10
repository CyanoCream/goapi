package router

import (
	"challenge-08/controllers"
	"challenge-08/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("item")
	productRouter.Use(middlewares.Authentication())
	{
		productRouter.GET("/products", controllers.GetAllProducts)
		productRouter.GET("/products/:productID", controllers.GetProduct)
		productRouter.POST("/products", controllers.CreateProduct)
		productRouter.PUT("/products/:productID", middlewares.AdminMiddleware(), controllers.UpdateProduct)
		productRouter.DELETE("/products/:productID", middlewares.AdminMiddleware(), controllers.DeleteProduct)
	}

	return r
}
