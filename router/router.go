package router

import (
	newsRoutes "news_api/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	newsRoutes.SetNewsRoutes(app)
}
