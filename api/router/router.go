package router

import (
	"github.com/44t4nk1/twitter-word-like/api/controller"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	app.Get("/tweets/:user", controller.GetTweets)
}
