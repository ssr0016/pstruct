package middleware

import (
	"task-management-system/internal/rbac/permissions"

	"github.com/gofiber/fiber/v2"
)

func PermissionMiddleware(permissionSrv permissions.Service, requiredPermissions ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// userID, ok := c.Locals("userID").(int) // Assumes userID is an int
		// if !ok {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		// }

		permissions, err := permissionSrv.GetPermissions(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch permissions"})
		}

		for _, p := range permissions {
			for _, requiredPermission := range requiredPermissions {
				if p.Name == requiredPermission {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You do not have the necessary permissions to access this resource"})
	}
}
