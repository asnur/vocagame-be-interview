package middleware

import (
	"github.com/asnur/vocagame-be-interview/internal/usecase/shared"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware creates JWT authentication middleware
func AuthMiddleware(sharedUseCase shared.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from Authorization header
		token, err := utils.ExtractTokenFromHeader(c)
		if err != nil {
			return utils.Response(c, fiber.StatusUnauthorized, "Unauthorized", nil, err)
		}

		// Validate token and get claims
		claims, err := sharedUseCase.ValidateTokenAndGetClaims(c.UserContext(), token)
		if err != nil {
			return utils.Response(c, fiber.StatusUnauthorized, "Invalid token", nil, err)
		}

		// Store user information in context
		c.Locals("user_id", claims.Data.UserId)
		c.Locals("user_claims", claims)
		c.Locals("token", token)

		return c.Next()
	}
}

// OptionalAuthMiddleware creates optional JWT authentication middleware
// This middleware doesn't return error if token is missing, but validates if present
func OptionalAuthMiddleware(sharedUseCase shared.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Try to extract token from Authorization header
		token, err := utils.ExtractTokenFromHeader(c)
		if err != nil {
			// Token is missing or invalid format, but that's OK for optional auth
			return c.Next()
		}

		// Validate token and get claims
		claims, err := sharedUseCase.ValidateTokenAndGetClaims(c.UserContext(), token)
		if err != nil {
			// Token exists but is invalid, that's an error
			return utils.Response(c, fiber.StatusUnauthorized, "Invalid token", nil, err)
		}

		// Store user information in context
		c.Locals("user_id", claims.Data.UserId)
		c.Locals("user_claims", claims)
		c.Locals("token", token)

		return c.Next()
	}
}

// GetUserID extracts user ID from fiber context
func GetUserID(c *fiber.Ctx) (int64, bool) {
	userID, ok := c.Locals("user_id").(int64)
	return userID, ok
}

// GetUserClaims extracts user claims from fiber context
func GetUserClaims(c *fiber.Ctx) (interface{}, bool) {
	claims := c.Locals("user_claims")
	return claims, claims != nil
}

// GetToken extracts token from fiber context
func GetToken(c *fiber.Ctx) (string, bool) {
	token, ok := c.Locals("token").(string)
	return token, ok
}
