package handlers

import (
	"database/sql"
	"practice-go-books/internal/models"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	db *sql.DB
}

func NewBookHandler(db *sql.DB) *BookHandler {
	return &BookHandler{
		db: db,
	}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book

	rows, err := h.db.Query("Select * from books")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.UUID, &book.Title, &book.Author)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	c.JSON(200, books)
}
