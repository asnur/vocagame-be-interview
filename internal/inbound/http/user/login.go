package user

import (
	ibModel "github.com/asnur/vocagame-be-interview/internal/inbound/http/model/user"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Login(fCtx *fiber.Ctx) error {
	var req ibModel.LoginRequest

	if err := utils.ParseAndValidate(fCtx, &req); err != nil {
		return utils.Response(fCtx, fiber.StatusBadRequest, "Invalid request", nil, err)
	}

	loginResponse, err := c.UseCase.User.Login(fCtx.UserContext(), req.ToUcModel())
	if err != nil {
		return utils.Response(fCtx, fiber.StatusInternalServerError, "Failed to login user", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusOK, "User logged in successfully", loginResponse, nil)
}
