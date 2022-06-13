package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-booksapi/dtos"
	"go-booksapi/repositories"
)

func CreateBook(c *gin.Context) {
	var request dtos.BookRequestBody

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := repositories.Create(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Status(http.StatusCreated)
}

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repositories.FindAll())
}

func GetBook(c *gin.Context) {
	book, err := repositories.FindById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})

		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var request dtos.BookRequestBody

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	book, err := repositories.Update(id, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := repositories.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Book " + id + " deleted",
	})
}
