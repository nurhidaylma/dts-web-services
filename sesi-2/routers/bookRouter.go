package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-web-services/controllers"
)

func StartServe() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.GET("/book/:bookID", controllers.GetBookByID)
	router.GET("/books", controllers.GetBooks)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
