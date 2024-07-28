package main

import (
	"log"

	"github.com/aebalz/go-gin-starter/config"
	"github.com/aebalz/go-gin-starter/internal/book"
	"github.com/aebalz/go-gin-starter/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db, err := db.InitializePostgresDatabase()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Set up the Gin router
	r := gin.Default()

	// Initialize services and routes
	book.Init(db, r)

	// Run the server
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
