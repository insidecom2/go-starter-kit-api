package repositories

import (
	"go/starter-kit/api/src/consts"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryStruct struct {
	DB *sqlx.DB
}

type UserRepository interface {
	GetUserById(id float64) (r consts.UserResponse, err error)
}

func NewUserRepository(DB *sqlx.DB) *UserRepositoryStruct {
	return &UserRepositoryStruct{
		DB: DB,
	}
}
