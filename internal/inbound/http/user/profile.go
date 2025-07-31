package user

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/asnur/vocagame-be-interview/pkg/middleware"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetProfile(fCtx *fiber.Ctx) error {
	// Get user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserID(fCtx)
	if !ok {
		return utils.Response(fCtx, fiber.StatusUnauthorized, "User not authenticated", nil, nil)
	}

	// Here you would typically get user profile from database
	profile, err := c.UseCase.User.Profile(fCtx.UserContext(), ucModel.ProfileRequest{
		UserID: userID,
	})
	if err != nil {
		status, err := pkgErr.ErrorResPonse(err)
		return utils.Response(fCtx, status, "Failed to retrieve profile", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusOK, "Profile retrieved successfully", profile, nil)
}
