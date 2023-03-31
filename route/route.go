package route

import (
	"sesi_8/handler"
	"sesi_8/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/book")
	{
		api.POST("", server.CreateBook)
		api.GET("", server.GetBooks)
		api.GET("/:id", server.GetBook)
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)
	}
}
