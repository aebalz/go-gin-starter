package main

import (
	"fmt"

	"github.com/aebalz/go-gin-starter/config"
	"github.com/aebalz/go-gin-starter/internal/apis/book"
	"github.com/aebalz/go-gin-starter/middlewares/auth"
	"github.com/aebalz/go-gin-starter/pkg/db"
	"github.com/aebalz/go-gin-starter/pkg/server"
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
	auth.Init(db, r)

	// Run the server with graceful shutdown
	appPort := fmt.Sprintf(":%s", config.AppConfig.AppPort)
	server.RunServer(r, appPort)

}
