package routes

import (
	"ecom-backend/internal/server/data"
	"ecom-backend/internal/server/handler"
	AuthHandler "ecom-backend/internal/server/handler/auth"
	"ecom-backend/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

var app *fiber.App

func Start(a *fiber.App) {

	app = a

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/books/new", handler.AddBook)
	v1.Get("/books/all", handler.GetBooks)
	v1.Delete("/books/delete/:id", handler.DeleteBook)

	v1.Post("/user/new", AuthHandler.Signup)

	app.Get("/", templ.Handler(views.Page(views.Layout{
		Page: "home",
		Data: data.GetBooks(),
	})))
}
