package main

import (
	"github.com/44t4nk1/twitter-word-like/api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.MountRoutes(app)
	app.Listen(":3000")
}
