package main

import (
	"github.com/gin-gonic/gin"
	"pustaka-api/handler"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.RootHandler)
	r.GET("/ping", handler.PingHandler)
	r.GET("/query", handler.QueryHandler)

	// api grouping / versioning route
	api := r.Group("/api") // relative to localhost/api
	v1 := api.Group("/v1") // relative to localhost/api/v1
	v1.GET("/books/:id", handler.BookHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)

	v1.POST("/books", handler.PostBookHandler)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
