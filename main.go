package main

import (
	"fmt"
	"log"

	"github.com/aebalz/go-gin-starter/config"
	"github.com/aebalz/go-gin-starter/internal/book/handlers"
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

	// Set up the Gin router
	r := gin.Default()

	// Register routes
	handlers.RegisterBookRoutes(r, bookService)

	// Run the server
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
