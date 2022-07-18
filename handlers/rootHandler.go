package handlers

import "github.com/gofiber/fiber/v2"

func RootHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":    "welcome to simple rest api with go fiber",
		"author": "abdul15irsyad",
	})
}
