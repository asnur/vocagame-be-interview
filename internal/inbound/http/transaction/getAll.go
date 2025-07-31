package transaction

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/transaction"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/asnur/vocagame-be-interview/pkg/middleware"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAll(fCtx *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserID(fCtx)
	if !ok {
		return utils.Response(fCtx, fiber.StatusUnauthorized, "User not authenticated", nil, nil)
	}

	transactions, err := c.UseCase.Transaction.GetAll(fCtx.UserContext(), ucModel.TransactionGetAllRequest{
		UserID: userID,
	})
	if err != nil {
		status, err := pkgErr.ErrorResPonse(err)
		return utils.Response(fCtx, status, "Failed to get transactions", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusOK, "Transactions retrieved successfully", transactions, nil)
}
