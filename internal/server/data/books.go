package data

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"

	"github.com/charmbracelet/log"
)

func GetBooks() []models.Book {

	var books []models.Book

	err := initializer.DB.Find(&books).Error
	if err != nil {
		log.Error("Failure in books fetching")
	}

	if len(books) == 0 {
		log.Info("No Books found")
		return []models.Book{}
	}

	return books
}
