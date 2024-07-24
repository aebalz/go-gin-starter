package db

import (
	"fmt"
	"log"

	"github.com/aebalz/go-gin-starter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.AppConfig.PostgresUser, config.AppConfig.PostgresPassword, config.AppConfig.PostgresDb)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to PostgreSQL database!")
		return nil, err
	}

	return db, nil
}
