package controllers

import (
	"go-booksapi/config"
	"go-booksapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var conn *gorm.DB

func init() {
	conn = config.OpenConnection()
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	result := conn.Create(&book)

	c.IndentedJSON(http.StatusCreated, result.RowsAffected)
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book

	conn.Find(&books)

	c.IndentedJSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	conn.First(&book)
	// db.First(&book, "id = ?", id)

	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	conn.Delete(&book)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "id #" + id + " deleted",
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{ID: id}

	conn.First(&book)

	c.BindJSON(&book)

	conn.Save(&book)

	c.IndentedJSON(http.StatusOK, book)
}
