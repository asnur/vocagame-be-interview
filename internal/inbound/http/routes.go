package http

import (
	"github.com/asnur/vocagame-be-interview/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func (i Inbound) Routes(c *fiber.App) {
	// Register user routes
	service := c.Group("voca-wallets")
	v1 := service.Group("v1")

	// User routes
	user := v1.Group("user")
	user.Post("/register", i.User.Register)
	user.Post("/login", i.User.Login)

	// Protected routes - require authentication
	authMiddleware := middleware.AuthMiddleware(i.User.UseCase.Shared)
	user.Get("/profile", authMiddleware, i.User.GetProfile)

	// Wallet routes
	wallet := v1.Group("wallet", authMiddleware)
	wallet.Post("/", i.Wallet.CreateWallet)
	wallet.Post("/deposit", i.Wallet.Deposit)
	wallet.Post("/withdraw", i.Wallet.WithDrawl)
	wallet.Post("/transfer", i.Wallet.Transfer)
}
