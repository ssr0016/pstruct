// package middleware

// import (
// 	"task-management-system/pkg/util/jwt"

// 	"github.com/gofiber/fiber/v2"
// )

// func JWTProtected(secret string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		authHeader := c.Get("Authorization")
// 		if authHeader == "" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Missing or malformed JWT",
// 			})
// 		}

// 		tokenStr := authHeader[len("Bearer "):]

// 		claims, err := jwt.ValidateToken(tokenStr)
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid or expired JWT",
// 			})
// 		}

// 		// Set user ID in context
// 		c.Locals("userID", claims.UserID)

// 		// Set user roles in context
// 		c.Locals("userRoles", claims.Roles)

// 		return c.Next()
// 	}
// }

package middleware

import (
	"fmt"
	"strings"
	"task-management-system/pkg/util/jwt"

	"github.com/gofiber/fiber/v2"
)

// JWTProtected is a middleware to protect routes with JWT authentication.
func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		}

		// Ensure the Authorization header has the correct Bearer prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		tokenStr := authHeader[len("Bearer "):]

		// Validate the token and extract claims
		claims, err := jwt.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired JWT",
			})
		}

		// Debugging statements
		fmt.Printf("UserID from token: %s\n", claims.UserID)
		fmt.Printf("UserRoles from token: %v\n", claims.Roles)

		// Set user ID and roles in context
		c.Locals("userID", claims.UserID)
		c.Locals("userRoles", claims.Roles)

		return c.Next()
	}
}
