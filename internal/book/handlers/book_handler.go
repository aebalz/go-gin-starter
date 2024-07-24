package handlers

import (
	"github.com/aebalz/go-gin-starter/internal/book/services"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine, bookService services.BookService) {
	api := r.Group("/api/v1")
	{
		books := api.Group("/books")
		{
			books.GET("/", bookService.GetBooks)
			books.GET("/:id", bookService.GetBook)
			books.POST("/", bookService.CreateBook)
			books.PUT("/:id", bookService.UpdateBook)
			books.PATCH("/:id", bookService.PatchBook)
			books.DELETE("/:id", bookService.DeleteBook)
		}
	}
}
