package main

import (
	"fms/config"
	"fms/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/login", controller.Login)
}

func main() {
	app := fiber.New()
	config.InitDatabase()
	InitRoutes(app)

	app.Listen(":8080")
}
