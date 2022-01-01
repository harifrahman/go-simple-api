package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello there!",
		"creator": "Arif Rahman Hakim",
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput

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
