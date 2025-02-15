package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type RequestError struct {
	Errors []string
}

func (e RequestError) Error() string {
	return strings.Join(e.Errors, ", ")
}

func ParseAndValidateRequest[T interface{}](ctx *fiber.Ctx, request *T) (*T, error) {

	// Parse JSON request
	if err := ctx.BodyParser(request); err != nil {

		return request, err
	}

	// // Validate request
	if err := validate.Struct(request); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Validation failed: " + err.Tag()
		}
		// return nil, fiber.Map{
		// 	"errors": errors,
		// }
		return request, err
	}

	return request, nil
}
