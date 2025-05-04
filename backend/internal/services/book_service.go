package services

import (
	"github.com/DB-Vincent/want-to-read/internal/database"
	"github.com/DB-Vincent/want-to-read/internal/models"
)

type BookService struct{}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) ListBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (s *BookService) AddBook(book *models.Book) (*models.Book, error) {
	result := database.DB.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (s *BookService) UpdateBook(id int, book *models.Book) (*models.Book, error) {
	var existingBook models.Book
	if err := database.DB.First(&existingBook, id).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Model(&existingBook).Updates(book).Error; err != nil {
		return nil, err
	}

	return &existingBook, nil
}