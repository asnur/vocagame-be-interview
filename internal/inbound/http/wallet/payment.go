package wallet

import (
	ibModel "github.com/asnur/vocagame-be-interview/internal/inbound/http/model/wallet"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/asnur/vocagame-be-interview/pkg/middleware"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Payment(fCtx *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserID(fCtx)
	if !ok {
		return utils.Response(fCtx, fiber.StatusUnauthorized, "User not authenticated", nil, nil)
	}

	var req ibModel.PaymentRequest
	if err := utils.ParseAndValidate(fCtx, &req); err != nil {
		return utils.Response(fCtx, fiber.StatusBadRequest, "Invalid request", nil, err)
	}

	payment, err := c.UseCase.Wallet.Payment(fCtx.UserContext(), req.ToUcModel(userID))
	if err != nil {
		status, err := pkgErr.ErrorResPonse(err)
		return utils.Response(fCtx, status, "Failed to process payment", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusCreated, "Payment processed successfully", payment, nil)
}
