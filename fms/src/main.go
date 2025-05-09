package main

import (
	"fms/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.InitDatabase()

	app.Listen(":8080")
}
