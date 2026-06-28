package routes

import (
	"database/sql"
	"practice-go-books/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, db *sql.DB) {
	bookHandler := handlers.NewBookHandler(db)
	r.GET("/books", bookHandler.GetBooks)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
