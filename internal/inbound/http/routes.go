package http

import "github.com/gofiber/fiber/v2"

func (i Inbound) Routes(c *fiber.App) {
	// Register user routes
	service := c.Group("voca-wallets")
	v1 := service.Group("v1")

	// User routes
	user := v1.Group("user")
	user.Post("/register", i.User.Register)
	user.Post("/login", i.User.Login)
}
