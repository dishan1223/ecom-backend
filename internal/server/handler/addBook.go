package handler

import (
	"ecom-backend/internal/models"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v3"
)

func AddBook(c fiber.Ctx) error {
	body := new(models.Book)

	if err := c.Bind().JSON(body); err != nil {
		log.Error("Failed to parse body data")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body data",
		})
	}

	return nil
}
