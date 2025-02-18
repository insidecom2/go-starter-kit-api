package handlers

import (
	"go/starter-kit/api/src/consts"
	"go/starter-kit/api/src/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlerStruct) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(float64)
	user, err := h.Service.GetUserById(userId)

	if err != nil {
		return utils.ResponseError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseStruct[consts.UserResponse]{
		Message: "OK",
		Data:    &user,
	})

}
