package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var book = []Book{
	{ID: "1", Title: "Harry potter", Author: "JK Rolling"},
	{ID: "2", Title: "The amazing spider-man", Author: "Stan lee"},
	{ID: "3", Title: "Superman", Author: "DC"},
}

func getBook(c *gin.Context) {
	c.JSON(http.StatusOK, book)
}
func newBook(c *gin.Context) {
	var datas Book

	if err := c.ShouldBindJSON(&datas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	book = append(book, datas)
	c.JSON(http.StatusCreated, book)
}
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, a := range book {
		if a.ID == id {
			book = append(book[:i], book[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}
func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello wolrd",
		})
	})
	r.GET("/book", getBook)
	r.POST("/newbook", newBook)
	r.DELETE("/delete/:id", deleteBook)
	r.Run()
}
