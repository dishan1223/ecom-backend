package handler

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func AddBook(c fiber.Ctx) error {
	body := new(models.Book)

	if err := c.Bind().JSON(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid fields detected",
		})
	}

	res := initializer.DB.Create(body)

	if res.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Error creating book",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": "Book created successfully",
	})
}
