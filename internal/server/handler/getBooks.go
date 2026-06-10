package handler

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v3"
)

func GetBooks(c fiber.Ctx) error {

	var books []models.Book

	err := initializer.DB.Find(&books).Error
	if err != nil {
		log.Error("Failure in books fetching")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch book data",
		})
	}

	if len(books) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "No books found",
		})
	}

	return c.Status(http.StatusOK).JSON(books)
}
