package routers

import (
	"codecloud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupFileIORoutes(r *gin.Engine) {
	fileio := r.Group("/fileio")
	{
		fileio.GET("/get-files", controllers.GetFiles)
		fileio.GET("/get-file", controllers.GetFile)
	}
}
