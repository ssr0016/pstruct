package middleware

// import (
// 	"task-management-system/pkg/util/jwt"

// 	"github.com/gofiber/fiber/v2"
// )

// func RoleBasedAccess(roles ...string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		authHeader := c.Get("Authorization")
// 		tokenStr := authHeader[len("Bearer "):]

// 		claims, err := jwt.ValidateToken(tokenStr)
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid or expired JWT",
// 			})
// 		}

// 		// Retrieve user roles from the database
// 		userRoles, err := role.GetUserRoles(claims.UserID)
// 		if err != nil {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 				"error": "Error retrieving user roles",
// 			})
// 		}

// 		// Check if the user has one of the required roles
// 		for _, userRole := range userRoles {
// 			for _, role := range roles {
// 				if userRole == role {
// 					return c.Next()
// 				}
// 			}
// 		}

// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 			"error": "You do not have the required permissions",
// 		})
// 	}
// }
