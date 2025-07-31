package utils

import "github.com/gofiber/fiber/v2"

type (
	Responses struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Error   *Error      `json:"error,omitempty"`
	}

	Error struct {
		Message string `json:"message"`
	}
)

func Response(fCtx *fiber.Ctx, status int, message string, data interface{}, err error) error {
	response := Responses{
		Status:  status,
		Message: message,
		Data:    data,
	}

	if err != nil {
		response.Error = &Error{Message: err.Error()}
	}

	return fCtx.Status(status).JSON(response)
}
