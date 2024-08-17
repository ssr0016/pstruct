package middleware

import (
	"task-management-system/pkg/util/jwt"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		}

		tokenStr := authHeader[len("Bearer "):]

		claims, err := jwt.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired JWT",
			})
		}

		// Set user ID in context
		c.Locals("userID", claims.UserID)

		// Set user roles in context
		c.Locals("userRoles", claims.Roles)

		return c.Next()
	}
}
