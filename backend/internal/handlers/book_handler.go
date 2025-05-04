package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DB-Vincent/want-to-read/internal/services"
	"github.com/DB-Vincent/want-to-read/internal/models"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// @Summary     List all books
// @Description Get a list of all books in the system
// @Tags        books
// @Produce     json
// @Success     200 {array}  models.Book
// @Failure     500 {string} string
// @Router      /books [get]
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

// @Summary     Add book
// @Description Add a book into the database
// @Tags        books
// @Produce     json
// @Param		book body object true "Book to add"
// @Success     200 {object}  models.Book
// @Failure     500 {string} string
// @Router      /book [post]
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