package services

import (
	"github.com/DB-Vincent/want-to-read/internal/database"
	"github.com/DB-Vincent/want-to-read/internal/models"
)

type BookService struct{}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) ListBooks(userID uint) ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Where("user_id = ?", userID).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (s *BookService) AddBook(book *models.Book, userID uint) (*models.Book, error) {
	book.UserID = userID
	result := database.DB.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (s *BookService) UpdateBook(id int, userID uint, book *models.Book) (*models.Book, error) {
	var existingBook models.Book
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&existingBook).Error; err != nil {
		return nil, err
	}
	if err := database.DB.Model(&existingBook).Updates(book).Error; err != nil {
		return nil, err
	}
	return &existingBook, nil
}

func (s *BookService) DeleteBook(id int, userID uint) (int, error) {
	var existingBook models.Book
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&existingBook).Error; err != nil {
		return 0, err
	}
	if err := database.DB.Delete(&existingBook).Error; err != nil {
		return 0, err
	}
	return id, nil
}
