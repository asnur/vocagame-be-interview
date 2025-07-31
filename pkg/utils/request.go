package utils

import (
	"encoding/xml"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidate(c *fiber.Ctx, out interface{}) error {
	contentType := c.Get("Content-Type")

	switch {
	case contentType == fiber.MIMEApplicationJSON:
		if err := c.BodyParser(out); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
		}
	case contentType == fiber.MIMEApplicationXML, contentType == fiber.MIMETextXML:
		if err := xml.Unmarshal(c.Body(), out); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid XML body")
		}
	case contentType == fiber.MIMEApplicationForm, strings.HasPrefix(contentType, fiber.MIMEMultipartForm):
		if err := c.BodyParser(out); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid form body")
		}
	}

	// Query dan Path param parsing
	if err := c.QueryParser(out); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}
	if err := c.ParamsParser(out); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid path parameters")
	}

	// Validasi
	if err := validate.Struct(out); err != nil {
		// Bisa di-custom untuk return error field
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return nil
}
