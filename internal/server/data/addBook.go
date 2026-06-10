package data

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"

	"github.com/charmbracelet/log"
)

func AddBook(book models.Book) error {
	if err := initializer.DB.Create(&book).Error; err != nil {
		log.Error("Error Creating Book")
		return err
	}

	return nil
}
