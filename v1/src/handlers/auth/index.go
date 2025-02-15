package handlers

import (
	services "go/starter-kit/api/src/services/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlers interface {
	Register(ctx *fiber.Ctx) error
}

type AuthHandlersStruct struct {
	AuthService services.AuthServiceStruct
}

func NewAuthHandler(service services.AuthServiceStruct) *AuthHandlersStruct {
	return &AuthHandlersStruct{
		AuthService: service,
	}
}
