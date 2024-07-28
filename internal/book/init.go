package book

import (
	"fmt"

	"github.com/aebalz/go-gin-starter/internal/book/handlers"
	"github.com/aebalz/go-gin-starter/internal/book/models"
	"github.com/aebalz/go-gin-starter/internal/book/repositories"
	"github.com/aebalz/go-gin-starter/internal/book/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, r *gin.Engine) {
	// Auto Migrate
	fmt.Println("Migrate Data")
	db.AutoMigrate(&models.Author{}, &models.Book{})

	// Initialize Repository
	bookRepo := repositories.NewBookRepository(db)
	authRepo := repositories.NewAuthorRepository(db)

	// Initialize Service
	bookService := services.NewBookService(bookRepo)
	authService := services.NewAuthorService(authRepo)

	// Register Routes
	api := r.Group("/api/v1")
	handlers.RegisterBookRoutes(api, bookService)
	handlers.RegisterAuthorRoutes(api, authService)
}
