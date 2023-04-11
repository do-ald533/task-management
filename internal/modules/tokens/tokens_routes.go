package tokens

import (
	"github.com/gofiber/fiber/v2"
)

// All public routes of tokens module
func PublicRoutes(route fiber.Router) {
	route.Post("/tokens", CreateToken)
}
