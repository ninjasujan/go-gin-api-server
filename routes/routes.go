package routes

import (
	"ginserver/controller"

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
	// Book routes
	router.GET("/books", controller.GetBooks)
	router.POST("/books", controller.CreateBook)
	router.DELETE("/books/:id", controller.DeleteBook)
}
