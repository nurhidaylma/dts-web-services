package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-web-services/controllers"
	docs "github.com/nurhidaylma/dts-web-services/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Books API
// @version 1.0
// @description This is a sample service for managing books
// @termsOfService http://swagger.io/terms/
// @contact.name Ilma
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
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

	return r
}
