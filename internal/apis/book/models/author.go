package models

import (
	"github.com/aebalz/go-gin-starter/internal/apis"
	"github.com/aebalz/go-gin-starter/pkg/validate"
	"github.com/aebalz/go-gin-starter/utils"
)

// Book represents a book model
type Author struct {
	utils.CustomModel
	Name  string `gorm:"size:255;not null" validate:"required" json:"name"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}

// Implement the TableName method for Author
func (Author) TableName() string {
	return apis.AppName + "_" + "author"
}

// Validate method for Author
func (a *Author) Validate() error {
	return validate.ValidateStruct(a)
}
