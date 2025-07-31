package wallet

import (
	ibModel "github.com/asnur/vocagame-be-interview/internal/inbound/http/model/wallet"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/asnur/vocagame-be-interview/pkg/middleware"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) CreateWallet(fCtx *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserID(fCtx)
	if !ok {
		return utils.Response(fCtx, fiber.StatusUnauthorized, "User not authenticated", nil, nil)
	}

	// Bind request body to CreateWalletRequest
	var req ibModel.CreateWalletRequest
	if err := utils.ParseAndValidate(fCtx, &req); err != nil {
		return utils.Response(fCtx, fiber.StatusBadRequest, "Invalid request", nil, err)
	}

	// Convert to use case model and call use case
	wallet, err := c.UseCase.Wallet.CreateWallet(fCtx.UserContext(), req.ToUcModel(userID))
	if err != nil {
		status, err := pkgErr.ErrorResPonse(err)
		return utils.Response(fCtx, status, "Failed to create wallet", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusCreated, "Wallet created successfully", wallet, nil)
}
