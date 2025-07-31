package utils

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// formatValidationError mengubah pesan error validasi menjadi lebih user-friendly
func formatValidationError(err error) string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			tag := fieldError.Tag()
			param := fieldError.Param()

			var message string
			switch tag {
			case "required":
				message = fmt.Sprintf("%s is required", field)
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", field)
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters", field, param)
			case "max":
				message = fmt.Sprintf("%s must be at most %s characters", field, param)
			case "len":
				message = fmt.Sprintf("%s must be exactly %s characters", field, param)
			case "gt":
				message = fmt.Sprintf("%s must be greater than %s", field, param)
			case "gte":
				message = fmt.Sprintf("%s must be greater than or equal to %s", field, param)
			case "lt":
				message = fmt.Sprintf("%s must be less than %s", field, param)
			case "lte":
				message = fmt.Sprintf("%s must be less than or equal to %s", field, param)
			case "oneof":
				message = fmt.Sprintf("%s must be one of: %s", field, param)
			case "uuid":
				message = fmt.Sprintf("%s must be a valid UUID", field)
			case "alphanum":
				message = fmt.Sprintf("%s must contain only alphanumeric characters", field)
			case "numeric":
				message = fmt.Sprintf("%s must be numeric", field)
			case "url":
				message = fmt.Sprintf("%s must be a valid URL", field)
			case "datetime":
				message = fmt.Sprintf("%s must be a valid datetime format", field)
			default:
				message = fmt.Sprintf("%s is invalid", field)
			}

			errors = append(errors, message)
		}
	} else {
		// Fallback jika bukan ValidationErrors
		return err.Error()
	}

	if len(errors) == 1 {
		return errors[0]
	}

	return "Validation failed: " + strings.Join(errors, ", ")
}

// ValidationError represents a structured validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Tag     string `json:"tag"`
	Value   string `json:"value,omitempty"`
}

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
		return fiber.NewError(fiber.StatusUnprocessableEntity, formatValidationError(err))
	}

	return nil
}
