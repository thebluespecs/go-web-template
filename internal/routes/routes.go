package routes

import (
	"go-web-template/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	app.Get("/", controller.Health)
}
