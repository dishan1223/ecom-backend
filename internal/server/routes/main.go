package routes

import (
	"ecom-backend/internal/server/handler"

	"github.com/gofiber/fiber/v3"
)

var app *fiber.App

func Start(a *fiber.App) {
	app = a

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", handler.Hello)
	v1.Post("/add-book", handler.AddBook)
}
