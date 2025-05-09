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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user doesn't exists signup please!"})
	}
	token, err := utils.CreateToken("abcd")
	if err != nil {
		fmt.Println("failed to create token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})

}

func SignUp(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	//if user does't exist add new one

	var newUser models.Users

	result := config.DB.Where("email = ?", user.Email).First(&newUser)
	if result.Error != nil {
		newUser.Id = utils.GenerateID()
		newUser.Email = user.Email
		newUser.Name = user.Name
		newUser.Password = user.Password
		res := config.DB.Create(newUser)
		if res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to add new user!"})
		}
		//create token if user added

		token, err := utils.CreateToken(newUser.Id)
		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create jwt token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user already exists please login"})
}
