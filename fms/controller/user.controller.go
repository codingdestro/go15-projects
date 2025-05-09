package controller

import (
	"fms/config"
	"fms/models"
	"fms/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var myUser models.Users

	res := config.DB.Where("email = ?", user.Email).First(&myUser)

	if res.Error != nil {
		token, err := utils.CreateToken("abcd")
		if err != nil {
			fmt.Println("failed to create token")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	}

	return c.SendStatus(500)
}
