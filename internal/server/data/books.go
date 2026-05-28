package data

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"

	"github.com/charmbracelet/log"
)

func GetBooks() []models.Book {

	var books []models.Book

	res := initializer.DB.Find(&books)

	if res.Error != nil {
		log.Error("Failure in books fetching")
	}

	return books
}
