package main

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/Anjani100/go-login-api/models"
)

func main() {
	r := gin.Default() // Initializing a gin router

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	// })

	models.ConnectDatabase()

	r.Run()
}