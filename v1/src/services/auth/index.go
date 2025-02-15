package services

import (
	"go/starter-kit/api/src/consts"
	repositories "go/starter-kit/api/src/repositories/auth"
)

type AuthServiceStruct struct {
	AuthRepository repositories.AuthRepositoryStruct
}

type AuthService interface {
	Register(req consts.RegisterRequest) (r consts.RegisterResponse, err error)
}

func NewAuthService(repository repositories.AuthRepositoryStruct) *AuthServiceStruct {
	return &AuthServiceStruct{
		AuthRepository: repository,
	}
}
