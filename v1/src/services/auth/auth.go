package services

import (
	"go/starter-kit/api/src/consts"
	"go/starter-kit/api/src/utils"
)

func (s *AuthServiceStruct) Register(req consts.RegisterRequest) (r consts.RegisterResponse, err error) {
	req.Password, _ = utils.HashPassword(req.Password)
	return s.AuthRepository.Register(req)
}
