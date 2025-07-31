package transaction

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/transaction"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/asnur/vocagame-be-interview/pkg/middleware"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Get(fCtx *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserID(fCtx)
	if !ok {
		return utils.Response(fCtx, fiber.StatusUnauthorized, "User not authenticated", nil, nil)
	}

	transactionID := fCtx.Params("id")
	if transactionID == "" {
		return utils.Response(fCtx, fiber.StatusBadRequest, "Transaction ID is required", nil, nil)
	}

	transaction, err := c.UseCase.Transaction.Get(fCtx.UserContext(), ucModel.TransactionGet{
		UserID: userID,
		TrxID:  transactionID,
	})
	if err != nil {
		status, err := pkgErr.ErrorResPonse(err)
		return utils.Response(fCtx, status, "Failed to get transaction", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusOK, "Transaction retrieved successfully", transaction, nil)
}
