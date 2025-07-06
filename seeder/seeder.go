package seeder

import "ginserver/controller"

// SeedBooks initializes the books data
func SeedBooks() {
	controller.Books = []controller.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", Price: 100},
		{ID: "2", Title: "Book 2", Author: "Author 2", Price: 200},
		{ID: "3", Title: "Book 3", Author: "Author 3", Price: 300},
	}
}
