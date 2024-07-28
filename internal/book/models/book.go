package models

import (
	"github.com/aebalz/go-gin-starter/pkg/validate"
	"gorm.io/gorm"
)

// Book represents a book model
type Book struct {
	gorm.Model
	Title    string `gorm:"size:255;not null" validate:"required"`
	AuthorID uint   `gorm:"not null" validate:"required"`
	// Author      Author  `gorm:"foreignKey:AuthorID"`
	Publisher   string  `gorm:"size:255" validate:"required"`
	PublishedAt string  `gorm:"size:255" validate:"required,datetime"`
	ISBN        string  `gorm:"size:20" validate:"required,len=13,numeric"`
	Price       float64 `gorm:"type:decimal(10,2)" validate:"required,gt=0"`
}

// Validate method for Book
func (b *Book) Validate() error {
	return validate.ValidateStruct(b)
}
