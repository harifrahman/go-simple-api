package book

import "encoding/json"

/*
	use something like below to specify input name
	Author string `json:"authorName"`
*/
type BookInput struct {
	Title  string      `json:"title" binding:"required"`
	Price  json.Number `json:"price" binding:"required,number"`
	Author string      `json:"authorName"`
}
