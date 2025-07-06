package main

import (
	"fmt"
	"ginserver/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.RegisterRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
