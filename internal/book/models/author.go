package models

import (
	"github.com/aebalz/go-gin-starter/pkg/validate"
	"gorm.io/gorm"
)

// Book represents a book model
type Author struct {
	gorm.Model
	Name  string `gorm:"size:255;not null" validate:"required"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}

// Validate method for Author
func (a *Author) Validate() error {
	return validate.ValidateStruct(a)
}
