package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(route fiber.Router) {
	route.Get("/healthcheck", GetHealthcheck)
}
