package models

import (
	"gorm.io/gorm"
)

// Book represents a book model
type Book struct {
	gorm.Model
	Title       string
	Author      string
	Publisher   string
	PublishedAt string
	ISBN        string
	Price       float64
}
