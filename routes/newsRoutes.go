package newsRoutes

import (
	newsHandler "news_api/handlers/news"

	"github.com/gofiber/fiber/v2"
)

func SetNewsRoutes(router fiber.Router) {
	news := router.Group("/news")

	news.Get("/list", newsHandler.GetNews)
	news.Post("/edit/:Id", newsHandler.PostNews)
}
