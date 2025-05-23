package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DB-Vincent/want-to-read/internal/database"
	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/DB-Vincent/want-to-read/internal/services"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// @Summary		List all books
// @Description	Get a list of all books in the system
// @Tags			books
// @Produce		json
// @Success		200	{array}		models.Book
// @Failure		500	{string}	string
// @Router			/api/books [get]
func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.ListBooks()
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// @Summary		Add book
// @Description	Add a book into the database
// @Tags			books
// @Produce		json
// @Param			book	body		object	true	"Book to add"
// @Success		200		{object}	models.Book
// @Failure		500		{string}	string
// @Router			/api/book [post]
func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdBook, err := h.bookService.AddBook(&book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createdBook); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// @Summary		Update book
// @Description	Updates a book based on the given ID
// @Tags			books
// @Produce		json
// @Param			id		path		int		true	"ID of book"
// @Param			book	body		object	true	"Adjusted book object"
// @Success		200		{object}	models.Book
// @Failure		500		{string}	string
// @Router			/api/book/{id} [patch]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid book ID",
		})
		return
	}

	// Get the existing book first
	var existingBook models.Book
	if err := database.DB.First(&existingBook, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	// Parse the update request
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Update only the fields that were provided
	if err := database.DB.Model(&existingBook).Updates(updateData).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update book",
		})
		return
	}

	c.JSON(200, existingBook)
}

// @Summary		Delete book
// @Description	Deletes a book based on the given ID
// @Tags			books
// @Produce		json
// @Param			id		path		int		true	"ID of book"
// @Success		200		{string}	string
// @Failure		500		{string}	string
// @Router			/api/book/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid book ID",
		})
		return
	}

	deletedBook, err := h.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete book",
		})
		return
	}

	c.JSON(200, deletedBook)
}
