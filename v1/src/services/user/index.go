package services

import (
	"go/starter-kit/api/src/consts"
	repositories "go/starter-kit/api/src/repositories/user"
)

type UserServiceStruct struct {
	Repo repositories.UserRepositoryStruct
}
type UserService interface {
	GetUserById(id float64) (r consts.UserResponse, err error)
}

func NewUserService(repo repositories.UserRepositoryStruct) *UserServiceStruct {
	return &UserServiceStruct{
		Repo: repo,
	}
}
