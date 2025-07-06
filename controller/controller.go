package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var Books = []Book{}

// GetBooks returns all books
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, Books)
}

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Books = append(Books, book)
	c.JSON(http.StatusCreated, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
