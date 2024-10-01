package controllers

import (
	"codecloud/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var fileIOService = service.NewFileIOService()

func GetFiles(c *gin.Context) {
	startingPath, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current directory"})
		return
	}

	data, err := fileIOService.ParseAndGetFolders(startingPath, startingPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File path query parameter is required"})
		return
	}

	startingPath, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current directory"})
		return
	}

	fileData, err := fileIOService.GetFileData(startingPath, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": string(fileData)})
}
