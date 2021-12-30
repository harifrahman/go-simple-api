package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", rootHandler)

	r.GET("/ping", pingHandler)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello there!",
		"creator": "Arif Rahman Hakim",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
