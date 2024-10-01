package controllers

import (
	"codecloud/service"

	"github.com/gin-gonic/gin"
)

var wsService = service.NewWSService()

func WS(c *gin.Context) {
	wsService.ConnectWSAndListen(c)
}
