package middleware

import (
	"fms/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthToken(c *fiber.Ctx) error {

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}
	c.Locals("userId", claims["userId"])
	return c.Next()
}
