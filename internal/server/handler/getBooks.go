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

	res := initializer.DB.Find(&books)

	if res.Error != nil {
		log.Error("Failure in books fetching")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch book data",
		})
	}

	return c.Status(http.StatusOK).JSON(books)
}
