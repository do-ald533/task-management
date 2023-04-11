package middleware

import (
	"goapi/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func Credentials(permissons ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, _ := jwt.ExtractTokenMetadata(c)

		for _, has_permission := range claims.Credentials {
			// Only user with provided role will pass.
			if !has_permission {
				// Return status 403 and permission denied error message.
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"role":               claims.Role,
					"needed_permissions": permissons,
					"error":              true,
					"msg":                "Permission denied, check role and credentials of your token",
				})
			}
		}

		return c.Next()
	}
}

func CredentialsByRole(roles ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, _ := jwt.ExtractTokenMetadata(c)

		for _, has_roles := range roles {
			// Only user with provided role will pass.
			if has_roles != claims.Role {
				// Return status 403 and permission denied error message.
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"role":         claims.Role,
					"needed_roles": roles,
					"error":        true,
					"msg":          "Role denied, check the role of your token",
				})
			}
		}

		return c.Next()
	}
}
