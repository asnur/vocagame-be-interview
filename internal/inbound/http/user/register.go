package user

import (
	ibModel "github.com/asnur/vocagame-be-interview/internal/inbound/http/model/user"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/asnur/vocagame-be-interview/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Register(fCtx *fiber.Ctx) error {
	var req ibModel.RegisterRequest

	if err := utils.ParseAndValidate(fCtx, &req); err != nil {
		return utils.Response(fCtx, fiber.StatusBadRequest, "Invalid request", nil, err)
	}

	user, err := c.UseCase.User.Register(fCtx.UserContext(), req.ToUcModel())
	if err != nil {
		status, err := pkgErr.ErrorResPonse(err)

		return utils.Response(fCtx, status, "Failed to register user", nil, err)
	}

	return utils.Response(fCtx, fiber.StatusCreated, "User registered successfully", user, nil)
}
