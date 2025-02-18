package handlers

import (
	services "go/starter-kit/api/src/services/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandlerStruct struct {
	Service services.UserServiceStruct
}

type UserHandler interface {
	GetUserById(ctx *fiber.Ctx) error
}

func NewUserHandler(service services.UserServiceStruct) *UserHandlerStruct {
	return &UserHandlerStruct{
		Service: service,
	}
}
