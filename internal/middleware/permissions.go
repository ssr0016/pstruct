package middleware

import (
	"context"
	"fmt"
	"task-management-system/internal/rbac/permissions/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

// PermissionMiddleware is a middleware to check if the user has the required permissions
func PermissionMiddleware(action string, repo *postgres.PermissionsRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hasPermission, err := checkUserPermission(c.Context(), repo, action)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No permission"})
		}

		return c.Next()
	}
}
func checkUserPermission(ctx context.Context, repo *postgres.PermissionsRepository, action string) (bool, error) {
	permissions, err := repo.GetUserPermissions(ctx)
	if err != nil {
		return false, fmt.Errorf("error retrieving user permissions: %v", err)
	}

	for _, perm := range permissions {
		if perm.Name == action {
			return true, nil
		}
	}

	return false, nil
}
