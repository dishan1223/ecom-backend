package handler

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func DeleteBook(c fiber.Ctx) error {
	/*
		if c.IP() != ClientIP {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "This api is only allowed for specific ip's",
			})
		}
	*/

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Book ID is required",
		})
	}

	res := initializer.DB.Delete(&models.Book{}, id)
	if res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete your book",
		})
	}

	if res.RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "No book found with given Book ID",
		})
	}

	return c.SendStatus(http.StatusOK)
}
