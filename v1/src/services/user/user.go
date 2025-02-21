package services

import (
	"go/starter-kit/api/src/consts"
)

func (s *UserServiceStruct) GetUserById(id float64) (r consts.UserResponse, err error) {

	r, err = s.Repo.GetUserById(id)
	if err != nil {
		return r, err
	}

	return r, nil

}
