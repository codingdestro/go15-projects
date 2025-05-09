package main

import (
	"fms/config"
	"fms/controller"
	"fms/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/login", controller.Login)
	app.Post("/signup", controller.SignUp)
	app.Use("/user", middleware.AuthToken)
	app.Post("/user", home)
	app.Post("/user/folders", controller.FetchFolders)
	app.Post("/user/folders/create/:name", controller.CreateFolder)
	app.Delete("/user/folders/:folderID", controller.DeleteFolder)
}

func main() {
	app := fiber.New()
	config.InitDatabase()
	InitRoutes(app)

	app.Listen(":8080")
}

func home(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	fmt.Println(userId)

	return c.SendStatus(200)

}
