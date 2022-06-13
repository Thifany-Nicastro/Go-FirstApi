package controllers

import (
	"go-booksapi/config"
	"go-booksapi/dtos"
	"go-booksapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var data dtos.BookRequestBody

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:  data.Title,
		Author: data.Author,
	}

	if result := config.Database.Create(&book); result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book

	config.Database.Find(&books)

	c.IndentedJSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	// var param dtos.BookRequestParams

	// if err := c.ShouldBindUri(&param); err != nil {
	// 	c.JSON(400, gin.H{
	// 		"msg": err.Error(),
	// 	})
	// 	return
	// }

	// book := models.Book{
	// 	ID: param.ID,
	// }

	id := c.Param("id")

	book := models.Book{ID: id}

	if result := config.Database.First(&book); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})

		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	if result := config.Database.Delete(&book); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})

		return
	}

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
