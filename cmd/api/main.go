package main

import (
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/server/handler"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
	initializer.SyncDB()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	PORT := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/", handler.Hello)

	log.Fatal(app.Listen(":" + PORT))
}
