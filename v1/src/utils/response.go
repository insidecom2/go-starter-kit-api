package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

const (
	ErrorCodeUniqueViolation     = "23505"
	ErrorCodeForeignKeyViolation = "23503"
	ErrorCodeNotNullViolation    = "23502"

	ErrorMessageConflict       = "Conflict"
	ErrorMessageBadRequest     = "Bad request"
	ErrorMessageUnauthorized   = "Unauthorized"
	ErrorMessageNotfound       = "Not found"
	ErrorMessageForbidden      = "Forbidden"
	ErrorMessageInternalServer = "Internal server error"
)

type ResponseStruct[T interface{}] struct {
	Message string `json:"message"`
	Data    *T     `json:"data"`
	Error   *error `json:"error"`
}

func ResponseError(ctx *fiber.Ctx, err error) error {

	var errMessage interface{}
	errMessage = err

	// validator handler error
	_, isValidator := err.(validator.ValidationErrors)
	if isValidator {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Validation failed: " + err.Tag()
		}
		errMessage = fiber.Map{
			"errors": errors,
		}
	}

	// handler error postgres
	if pgErr, ok := err.(*pq.Error); ok {
		errMessage = fiber.Map{
			"errors": pgErr.Detail,
		}

	}
	return ctx.Status(fiber.StatusBadRequest).JSON(errMessage)

}
