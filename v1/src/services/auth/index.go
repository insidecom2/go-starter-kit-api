package services

import (
	"go/starter-kit/api/src/configs"
	"go/starter-kit/api/src/consts"
	repositories "go/starter-kit/api/src/repositories/auth"
)

type AuthServiceStruct struct {
	AuthRepository repositories.AuthRepositoryStruct
	config         *configs.Config
}

type AuthService interface {
	Register(req consts.RegisterRequest) (r consts.UserResponse, err error)
	Login(req consts.LoginRequest) (r consts.LoginResponse, err error)
}

func NewAuthService(repository repositories.AuthRepositoryStruct, config *configs.Config) *AuthServiceStruct {
	return &AuthServiceStruct{
		AuthRepository: repository,
		config:         config,
	}
}
