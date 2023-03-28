package routers

import (
	"Latihan1/controllers"
	"github.com/gin-gonic/gin"
)

func BookRouter() {
	var err error
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
