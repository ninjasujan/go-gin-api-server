package main

import (
	"fmt"
	"ginserver/routes"
	"ginserver/seeder"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Seed initial data
	seeder.SeedBooks()

	routes.RegisterRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
