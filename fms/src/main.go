package main

import (
	"fms/config"
	"fms/controller"
	"fms/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Static("/", "./public")

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("./static/login.html")
	})
	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendFile("./static/signup.html")
	})
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendFile("./static/dashboard.html")
	})

	app.Post("/api/auth/login", controller.Login)
	app.Post("/api/auth/signup", controller.SignUp)
	app.Use("/api/user", middleware.AuthToken)
	app.Post("/api/user/folders", controller.FetchFolders)
	app.Post("/api/user/folders/create/:name", controller.CreateFolder)
	app.Delete("/api/user/folders/:folderID", controller.DeleteFolder)

	app.Post("/api/user/upload", controller.Upload)
	app.Get("/api/user/files/:folderId", controller.GetAllFiles)

	app.Post("/user", home)
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
