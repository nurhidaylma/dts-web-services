package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-web-services/controllers"
	docs "github.com/nurhidaylma/dts-web-services/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServe() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/book", controllers.CreateBook)
		v1.PUT("/book/:bookID", controllers.UpdateBook)
		v1.GET("/book/:bookID", controllers.GetBookByID)
		v1.GET("/books", controllers.GetBooks)
		v1.DELETE("/book/:bookID", controllers.DeleteBook)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

	return r
}
