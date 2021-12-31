package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)
	r.GET("/ping", pingHandler)
	r.GET("/query", queryHandler)

	// api grouping / versioning route
	api := r.Group("/api") // relative to localhost/api
	v1 := api.Group("/v1") // relative to localhost/api/v1
	v1.GET("/books/:id", bookHandler)
	v1.GET("/books/:id/:title", bookHandler)

	v1.POST("/books", postBookHandler)

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

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Unprocessable Entity on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			// fmt.Println(err)
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errorMessages,
		})
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
	Title  string      `json:"title" binding:"required"`
	Price  json.Number `json:"price" binding:"required,number"`
	Author string      `json:"authorName"`
}
