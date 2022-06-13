package repositories

import (
	"errors"

	"go-booksapi/config"
	"go-booksapi/dtos"
	"go-booksapi/models"
)

func FindAll() []models.Book {
	var books []models.Book

	config.Database.Find(&books)

	return books
}

func FindById(id string) (models.Book, error) {
	book := models.Book{ID: id}

	if result := config.Database.First(&book); result.Error != nil {
		return models.Book{}, result.Error
	}

	return book, nil
}

func Create(request dtos.BookRequestBody) error {
	book := models.Book{
		Title:  request.Title,
		Author: request.Author,
	}

	if result := config.Database.Create(&book); result.Error != nil {
		return result.Error
	}

	return nil
}

func Update(id string, request dtos.BookRequestBody) (models.Book, error) {
	book := models.Book{ID: id}

	if result := config.Database.First(&book); result.Error != nil {
		return models.Book{}, result.Error
	}

	book.Title = request.Title
	book.Author = request.Author

	if result := config.Database.Save(&book); result.Error != nil {
		return models.Book{}, result.Error
	}

	return book, nil
}

func Delete(id string) error {
	book := models.Book{ID: id}

	if result := config.Database.Delete(&book); result.RowsAffected == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}
