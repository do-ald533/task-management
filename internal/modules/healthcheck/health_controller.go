package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

type success struct {
	Status int `json:"status"`
}

// @Description Health check test.
// @Summary Health check validation
// @Tags Healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} success
// @Router /healthcheck [get]
func GetHealthcheck(c *fiber.Ctx) error {
	jaum := c.Query("teste", "jaum")
	// Return status 200 OK
	teste := fiber.Map{"status": fiber.StatusOK, "teste": jaum}
	return c.Status(fiber.StatusOK).JSON(teste)
}
