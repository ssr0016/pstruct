package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RoleProtected(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, ok := c.Locals("userRoles").([]string)
		if !ok || len(roles) == 0 {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "No roles found in JWT",
			})
		}

		for _, userRole := range roles {
			for _, requiredRole := range requiredRoles {
				if userRole == requiredRole {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You do not have the necessary permissions to access this resource",
		})
	}
}
