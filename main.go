package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/read", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.POST("/event", func(context *gin.Context) {
		switch context.Request.FormValue("event") {
		case "impression":
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		case "click":
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		case "call":
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		default:
			context.JSON(http.StatusNotImplemented, gin.H{
				"message": "pong",
			})
		}
	})

	server.Run()
}
