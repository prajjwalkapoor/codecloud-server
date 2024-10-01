package routers

import (
	"codecloud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupWSRoutes(r *gin.Engine) {
	r.GET("/ws", controllers.WS)
}
