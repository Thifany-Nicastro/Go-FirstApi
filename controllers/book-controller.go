package controllers

import (
	"go-booksapi/config"
	"go-booksapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	result := config.Database.Create(&book)

	c.IndentedJSON(http.StatusCreated, result.RowsAffected)
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book

	config.Database.Find(&books)

	c.IndentedJSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	config.Database.First(&book)

	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	config.Database.Delete(&book)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "id #" + id + " deleted",
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	config.Database.First(&book)

	c.BindJSON(&book)

	config.Database.Save(&book)

	c.IndentedJSON(http.StatusOK, book)
}
