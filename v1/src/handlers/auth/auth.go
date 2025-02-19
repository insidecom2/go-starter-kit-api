package handlers

import (
	"go/starter-kit/api/src/consts"
	"go/starter-kit/api/src/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *authHandlersStruct) Register(ctx *fiber.Ctx) error {

	req, err := utils.ParseAndValidateRequest(ctx, &consts.RegisterRequest{})
	if err != nil {
		return utils.ResponseError(ctx, err)
	}

	res, err := h.AuthService.Register(*req)
	if err != nil {
		return utils.ResponseError(ctx, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ResponseStruct[consts.UserResponse]{
		Message: "Created",
		Data:    &res,
	})
}

func (h *authHandlersStruct) Login(ctx *fiber.Ctx) error {
	req, err := utils.ParseAndValidateRequest(ctx, &consts.LoginRequest{})
	if err != nil {
		return utils.ResponseError(ctx, err)
	}

	res, err := h.AuthService.Login(*req)
	if err != nil {
		return utils.ResponseError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseStruct[consts.LoginResponse]{
		Message: "Login Success",
		Data:    &res,
	})
}
