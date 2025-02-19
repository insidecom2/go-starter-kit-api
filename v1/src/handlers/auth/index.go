package handlers

import (
	services "go/starter-kit/api/src/services/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type authHandlersStruct struct {
	AuthService services.AuthService
}

func NewAuthHandler(service services.AuthService) AuthHandlers {
	return &authHandlersStruct{
		AuthService: service,
	}
}
