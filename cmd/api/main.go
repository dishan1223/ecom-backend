package main

import (
	"ecom-backend/internal/config"
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/server/routes"
	"ecom-backend/internal/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
	initializer.SyncDB()
	service.Init(os.Getenv("JWT_SECRET"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	PORT := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		StructValidator: config.NewValidator(),
	})

	routes.Start(app)

	log.Fatal(app.Listen(":" + PORT))
}
