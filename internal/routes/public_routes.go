package routes

import (
	"goapi/internal/modules/healthcheck"
	"goapi/internal/modules/tokens"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	api := app.Group("")
	healthcheck.PublicRoutes(api)
	tokens.PublicRoutes(api)
}
