package main

import (
	"ecom-backend/consts"
	"ecom-backend/internal/config"
	"ecom-backend/internal/initializer"
	"ecom-backend/internal/server/routes"
	"ecom-backend/internal/service"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func init() {
	initializer.MustLoadEnv()
	initializer.MustConnectToDB()
	initializer.SyncDB()
	service.Init(config.MustLoadEnv(consts.JWT_SECRET))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	PORT := config.MustLoadEnv(consts.PORT)

	app := fiber.New(fiber.Config{
		StructValidator: config.NewValidator(),
	})

	routes.Start(app)

	log.Fatal(app.Listen(":" + PORT))
}
