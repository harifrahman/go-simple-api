package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.GET("/ping", pingHandler)
	r.GET("/query", queryHandler)

	r.GET("/books/:id", bookHandler)
	r.GET("/books/:id/:title", bookHandler)

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

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
