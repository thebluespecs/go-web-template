package main

import (
	"go-web-template/config"
	"go-web-template/internal/logger"
	"go-web-template/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.Info("Starting server...")
	app := fiber.New()
	routes.AddRoutes(app)
	app.Listen(":" + config.Get("PORT"))
}
