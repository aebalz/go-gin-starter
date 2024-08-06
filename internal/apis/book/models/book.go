package models

import (
	"github.com/aebalz/go-gin-starter/internal/apis"
	"github.com/aebalz/go-gin-starter/pkg/validate"
	"github.com/aebalz/go-gin-starter/utils"
)

// Book represents a book model
type Book struct {
	utils.CustomModel
	Title       string  `gorm:"size:255;not null" validate:"required" json:"title"`
	AuthorID    uint    `gorm:"not null" validate:"required" json:"author_id"`
	Publisher   string  `gorm:"size:255" validate:"required" json:"publisher"`
	PublishedAt string  `gorm:"size:255" validate:"required,datetime" json:"published_at"`
	ISBN        string  `gorm:"size:20" validate:"required,len=13,numeric" json:"isbn"`
	Price       float64 `gorm:"type:decimal(10,2)" validate:"required,gt=0" json:"price"`
	// Author      Author  `gorm:"foreignKey:AuthorID"`
}

// Implement the TableName method for Author
func (Book) TableName() string {
	return apis.AppName + "_" + "book"
}

// Validate method for Book
func (b *Book) Validate() error {
	return validate.ValidateStruct(b)
}
