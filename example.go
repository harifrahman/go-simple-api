package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.GET("/ping", pingHandler)
	r.GET("/query", queryHandler)

	r.GET("/books/:id", bookHandler)
	r.GET("/books/:id/:title", bookHandler)

	r.POST("/books", postBookHandler)

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

func postBookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		// log.Fatal(err)
		c.JSON(http.StatusUnprocessableEntity, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"title":      bookInput.Title,
		"price":      bookInput.Price,
		"authorName": bookInput.Author,
	})
}

/*
	use something like below to specify input name
	Author string `json:"authorName"`
*/
type BookInput struct {
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required,number"`
	Author string `json:"authorName"`
}
