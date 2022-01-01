package main

import (
	"fmt"
	"log"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "pustaka:123456@tcp(127.0.0.1:3302)/pustaka_api?charset=utf8&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error!")
	}

	fmt.Println("Db connected successfully")

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
