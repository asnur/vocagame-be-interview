package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ExtractTokenFromHeader extracts Bearer token from Authorization header
func ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Authorization header is required")
	}

	// Check if the header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization header format")
	}

	// Extract the token (remove "Bearer " prefix)
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Token is required")
	}

	return token, nil
}

// ExtractTokenFromQuery extracts token from query parameter
func ExtractTokenFromQuery(c *fiber.Ctx, paramName string) (string, error) {
	token := c.Query(paramName)
	if token == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Token is required in query parameter")
	}

	return token, nil
}

// ExtractTokenFromCookie extracts token from cookie
func ExtractTokenFromCookie(c *fiber.Ctx, cookieName string) (string, error) {
	token := c.Cookies(cookieName)
	if token == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Token is required in cookie")
	}

	return token, nil
}
