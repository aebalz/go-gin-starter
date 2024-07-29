package book

import (
	"github.com/aebalz/go-gin-starter/internal/apis/book/handlers"
	"github.com/aebalz/go-gin-starter/internal/apis/book/models"
	"github.com/aebalz/go-gin-starter/internal/apis/book/repositories"
	"github.com/aebalz/go-gin-starter/internal/apis/book/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, r *gin.Engine) {
	// Auto Migrate
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
