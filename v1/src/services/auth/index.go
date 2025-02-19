package services

import (
	"go/starter-kit/api/src/configs"
	"go/starter-kit/api/src/consts"
	repositories "go/starter-kit/api/src/repositories/auth"
)

type authServiceStruct struct {
	AuthRepository repositories.AuthRepository
	config         *configs.Config
}

type AuthService interface {
	Register(req consts.RegisterRequest) (r consts.UserResponse, err error)
	Login(req consts.LoginRequest) (r consts.LoginResponse, err error)
}

func NewAuthService(repository repositories.AuthRepository, config *configs.Config) AuthService {
	return &authServiceStruct{
		AuthRepository: repository,
		config:         config,
	}
}
