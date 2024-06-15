package router

import (
	"github.com/DrLivsey00/ShopPraactice/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.HomeHandler)

	return app
}
