package middleware

import (
	"fmt"
	"strings"
	"task-management-system/internal/rbac/permissionuser/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

// // PermissionMiddleware is a middleware to check if the user has the required permissions
// func PermissionMiddleware(action string, repo *postgres.PermissionsRepository) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		hasPermission, err := checkUserPermission(c.Context(), repo, action)
// 		if err != nil {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
// 		}

// 		if !hasPermission {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No permission"})
// 		}

// 		return c.Next()
// 	}
// }
// func checkUserPermission(ctx context.Context, repo *postgres.PermissionsRepository, action string) (bool, error) {
// 	permissions, err := repo.GetUserPermissions(ctx)
// 	if err != nil {
// 		return false, fmt.Errorf("error retrieving user permissions: %v", err)
// 	}

// 	for _, perm := range permissions {
// 		if perm.Name == action {
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }

// func PermissionMiddleware(requiredAction string, repo *postgres.PermissionUser) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		userID, ok := c.Locals("userID").(int)
// 		if !ok {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "User ID not found in context",
// 			})
// 		}

// 		permissions, err := repo.GetAllUserPermissions(c.Context(), userID)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": fmt.Sprintf("Error retrieving permissions: %v", err),
// 			})
// 		}

// 		hasPermission := false
// 		for _, perm := range permissions {
// 			if perm.Action == requiredAction {
// 				hasPermission = true
// 				break
// 			}
// 		}

// 		if !hasPermission {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 				"error": "User does not have permission for this action",
// 			})
// 		}

//			return c.Next()
//		}
//	}

// func PermissionMiddleware(requiredAction string, repo *postgres.PermissionUser) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Retrieve userID from context as a string (email address)
// 		userID, ok := c.Locals("userID").(string)
// 		if !ok {
// 			fmt.Println("PermissionMiddleware: User ID not found in context")
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "User ID not found in context",
// 			})
// 		}

// 		// Log userID from context
// 		fmt.Printf("PermissionMiddleware: UserID=%s\n", userID)

// 		// Call the repository method with email userID
// 		permissions, err := repo.GetAllUserPermissions(c.Context(), userID)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": fmt.Sprintf("Error retrieving permissions: %v", err),
// 			})
// 		}

// 		hasPermission := false
// 		for _, perm := range permissions {
// 			// Split the actions by comma and check if the required action is present
// 			actions := strings.Split(perm.Action, ",")
// 			for _, action := range actions {
// 				fmt.Printf("PermissionMiddleware: Checking permission %s\n", action)
// 				if strings.TrimSpace(action) == requiredAction {
// 					hasPermission = true
// 					break
// 				}
// 			}
// 			if hasPermission {
// 				break
// 			}
// 		}

// 		if !hasPermission {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 				"error": "User does not have permission for this action",
// 			})
// 		}

// 		return c.Next()
// 	}
// }

func PermissionMiddleware(requiredAction string, repo *postgres.PermissionUserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(string)
		if !ok {
			fmt.Println("PermissionMiddleware: User ID not found in context")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID not found in context",
			})
		}

		// Log userID from context
		fmt.Printf("PermissionMiddleware: UserID=%s\n", userID)

		// Call repo to get permissions for the user
		permissions, err := repo.GetAllUserPermissions(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Error retrieving permissions: %v", err),
			})
		}

		// Check if the user has the required permission
		hasPermission := false
		for _, perm := range permissions {
			// Split the actions by comma and check if the required action is present
			actions := strings.Split(perm.Action, ",")
			for _, action := range actions {
				fmt.Printf("PermissionMiddleware: Checking permission %s\n", action)
				if strings.TrimSpace(action) == requiredAction {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "User does not have permission for this action",
			})
		}

		return c.Next()
	}
}
