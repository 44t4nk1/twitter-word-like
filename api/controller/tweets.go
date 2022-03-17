package controller

import "github.com/gofiber/fiber/v2"

func GetTweets(c *fiber.Ctx) error {
	username := c.Params("user", "")
	if username == "" {
		return c.Status(404).SendString("Not Found")
	}
	return c.Status(404).SendString(username)
}
