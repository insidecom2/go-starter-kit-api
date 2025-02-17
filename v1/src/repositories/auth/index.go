package repositories

import (
	"go/starter-kit/api/src/consts"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryStruct struct {
	DB *sqlx.DB
}

type AuthRepository interface {
	Register(req consts.RegisterRequest) (r consts.UserResponse, err error)
	Login(req consts.LoginRequest) (r consts.UserResponse, err error)
	UpdateRefreshToken(req consts.UserEntity) (err error)
}

func NewAuthRepository(DB *sqlx.DB) *AuthRepositoryStruct {
	return &AuthRepositoryStruct{
		DB: DB,
	}
}
