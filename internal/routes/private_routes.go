package routes

import (
	"goapi/internal/middleware"
	"goapi/internal/modules/books"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	api := app.Group("", middleware.JWTProtected())

	books.PrivateRoutes(api)
}
