package middleware

import (
	"fmt"
	"strings"
	"task-management-system/internal/rbac/userroles"
	"task-management-system/internal/rbac/userroles/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

// RoleBasedAccessControl restricts access based on exact role matching
func RoleBasedAccessControl(repo *postgres.UserRoleRepository, requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID not found in context",
			})
		}

		userRoles, err := repo.GetUserRolesByID(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get user roles",
			})
		}

		// Create a set of user roles for easy lookup
		roleSet := parseRoles(userRoles)

		// Debugging information
		fmt.Printf("UserID from token: %s\n", userID)
		fmt.Printf("User roles: %v\n", roleSet)
		fmt.Printf("Required roles: %v\n", requiredRoles)

		// Check if all required roles are present
		for _, requiredRole := range requiredRoles {
			if !roleSet[requiredRole] {
				fmt.Printf("Role not found: %s\n", requiredRole)
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Access denied: user does not have the required roles",
				})
			}
		}

		return c.Next()
	}
}

// parseRoles converts a slice of UserRole to a set of roles.
func parseRoles(userRoles []*userroles.UserRole) map[string]bool {
	roleSet := make(map[string]bool)
	for _, role := range userRoles {
		// Remove curly braces and split by comma
		roles := strings.Trim(role.RoleNames, "{}")
		roleArray := strings.Split(roles, ",")
		for _, r := range roleArray {
			roleName := strings.TrimSpace(r)
			if roleName != "" {
				roleSet[roleName] = true
			}
		}
	}
	return roleSet
}
