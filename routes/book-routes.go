package routes

import (
	"go-booksapi/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
