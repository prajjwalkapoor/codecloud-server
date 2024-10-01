package main

import (
	"codecloud/config"
	"codecloud/middleware"
	"codecloud/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.LoadConfig()

	r.Use(middleware.CorsMiddleware())

	routers.SetupFileIORoutes(r)
	routers.SetupWSRoutes(r)

	r.Run(config.AppConfig.Port)
}
