package handlers

import (
	"github.com/aebalz/go-gin-starter/internal/book/services"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.RouterGroup, bookService services.BookService) {
	books := r.Group("/books")
	{
		books.GET("/", bookService.GetBooks)
		books.GET("/:id", bookService.GetBook)
		books.POST("/", bookService.CreateBook)
		books.PUT("/:id", bookService.UpdateBook)
		books.PATCH("/:id", bookService.PatchBook)
		books.DELETE("/:id", bookService.DeleteBook)
	}
}
