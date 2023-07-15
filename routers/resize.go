package routers

import (
	"github.com/Kyutech-C3/img-resize-server/controllers"

	"github.com/gin-gonic/gin"
)

func Resize(e *gin.RouterGroup) {
	resize := e.Group("/resize")
	{
		resize.POST("/", controllers.PostResize)
		resize.GET("/", controllers.GetResize)
		resize.GET("/:id", controllers.GetResizeStatus)
	}
}
