package services

import "go/starter-kit/api/src/consts"

func (s *UserServiceStruct) GetUserById(id float64) (r consts.UserResponse, err error) {
	return s.Repo.GetUserById(id)
}
