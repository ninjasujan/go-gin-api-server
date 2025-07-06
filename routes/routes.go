package routes

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

var Books = []Book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Price: 100},
	{ID: "2", Title: "Book 2", Author: "Author 2", Price: 200},
	{ID: "3", Title: "Book 3", Author: "Author 3", Price: 300},
}

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(router *gin.Engine) {
	// Add your routes here
	router.GET("/books", func(c *gin.Context) {
		c.JSON(200, Books)
	})

	router.POST("/books", func(ctx *gin.Context) {
		var book Book
		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Books = append(Books, book)
		ctx.JSON(http.StatusCreated, book)
	})

	router.DELETE("/books/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i, book := range Books {
			if book.ID == id {
				Books = append(Books[:i], Books[i+1:]...)
				ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})
}
