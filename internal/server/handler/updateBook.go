package handler

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func UpdateBook(c fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Book id is required",
		})
	}

	payload := new(models.Book)

	if err := c.Bind().Body(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Data Format",
		})
	}

	book := new(models.Book)

	if err := initializer.DB.Find(&book, id); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	if err := initializer.DB.Model(&book).Updates(payload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update book",
		})
	}

	return c.SendStatus(http.StatusOK)
}
