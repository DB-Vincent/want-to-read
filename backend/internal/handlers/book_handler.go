package handlers

import (
	"net/http"
	"strconv"

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
// @Tags		books
// @Produce		json
// @Param		user_id	path	int		true	"ID of user"
// @Success		200	{array}		models.Book
// @Failure		500	{string}	string
// @Failure		400	{string}	string
// @Router		/api/users/{user_id}/books [get]
func (h *BookHandler) ListBooks(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid user_id", "details": err.Error()})
		return
	}
	books, err := h.bookService.ListBooks(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary		Add book
// @Description	Add a book into the database
// @Tags		books
// @Produce		json
// @Param		book	body		object	true	"Book to add"
// @Param		user_id	path		int		true	"ID of user"
// @Success		200		{object}	models.Book
// @Failure		500		{string}	string
// @Failure		400		{string}	string
// @Router		/api/users/{user_id}/book [post]
func (h *BookHandler) AddBook(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid user_id", "details": err.Error()})
		return
	}
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	createdBook, err := h.bookService.AddBook(&book, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdBook)
}

// @Summary		Update book
// @Description	Updates a book based on the given ID
// @Tags		books
// @Produce		json
// @Param		id		path		int		true	"ID of book"
// @Param		user_id	path		int		true	"ID of user"
// @Param		book	body		object	true	"Adjusted book object"
// @Success		200		{string}	string
// @Failure		400		{string}	string
// @Failure		500		{string}	string
// @Router		/api/users/{user_id}/book/{id} [patch]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid user_id", "details": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID", "details": err.Error()})
		return
	}
	var updateData models.Book
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	updatedBook, err := h.bookService.UpdateBook(id, uint(userID), &updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

// @Summary		Delete book
// @Description	Deletes a book based on the given ID
// @Tags		books
// @Produce		json
// @Param		id		path		int		true	"ID of book"
// @Param		user_id	path		int		true	"ID of user"
// @Success		200		{string}	string
// @Failure		400		{string}	string
// @Failure		500		{string}	string
// @Router		/api/users/{user_id}/book/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid user_id", "details": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID", "details": err.Error()})
		return
	}
	deletedID, err := h.bookService.DeleteBook(id, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted_id": deletedID})
}
