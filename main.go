package main

import (
	"go-booksapi/config"
	"go-booksapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	router := gin.Default()

	routes.Routes(router)

	router.Run("localhost:8080")
}
