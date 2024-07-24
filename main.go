package main

import (
	"fmt"
	"net/http"

	"github.com/aebalz/go-gin-starter/config"
	"github.com/aebalz/go-gin-starter/internal/book/models"
	"github.com/aebalz/go-gin-starter/internal/book/repositories"
	"github.com/aebalz/go-gin-starter/internal/book/services"
	"github.com/aebalz/go-gin-starter/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	fmt.Println(config.AppConfig)
	db, err := db.InitializePostgresDatabase()
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.Book{})
	bookRepo := repositories.NewBookRepository(db)

	// Initialize the services with the repository
	bookService := services.NewBookService(bookRepo)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")
	{
		books := api.Group("/books")
		{
			books.GET("/", bookService.GetBooks)
			books.GET("/:id", bookService.GetBook)
			books.POST("/", bookService.CreateBook)
			books.PUT("/:id", bookService.UpdateBook)
			books.PATCH("/:id", bookService.UpdateBook)
			books.DELETE("/:id", bookService.DeleteBook)
		}
	}

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
