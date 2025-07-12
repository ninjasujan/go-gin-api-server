package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "ginserver/database"
	bookModel "ginserver/models"
)

// GetBooks returns all books
func GetBooks(c *gin.Context) {
	var books []bookModel.Book
	db.Database.Find(&books)
	c.JSON(http.StatusOK, books)
}

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	var book bookModel.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the book in the database
	result := db.Database.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// Delete the book from the database
	result := db.Database.Delete(&bookModel.Book{}, "id = ?", id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
