package services

import (
	"encoding/json"
	"go/starter-kit/api/src/consts"
	producer "go/starter-kit/api/src/producer/kafka"
	"go/starter-kit/api/src/utils"
)

func (s *authServiceStruct) Register(req consts.RegisterRequest) (r consts.UserResponse, err error) {
	req.Password, _ = utils.HashPassword(req.Password)
	return s.AuthRepository.Register(req)
}

func (s *authServiceStruct) Login(req consts.LoginRequest) (r consts.LoginResponse, err error) {
	user, err := s.AuthRepository.Login(req)

	if err != nil {
		return r, err
	}
	jwtToken := utils.NewJwtToken(s.config.JwtToken.Token)
	token, err := jwtToken.HashJwtToken(user)
	if err != nil {
		return r, err
	}
	reFreshToken, err := jwtToken.HashJwtToken(user)
	if err != nil {
		return r, err
	}
	_, err = s.AuthRepository.UpdateRefreshToken(consts.UserEntity{
		Id:           user.Id,
		ReFreshToken: &reFreshToken,
	})
	if err != nil {
		return r, err
	}

	r = consts.LoginResponse{
		UserResponse: user,
		Token:        token,
		ReFreshToken: reFreshToken,
	}

	msg, errs := json.Marshal(r)
	if errs != nil {
		return r, err
	}

	if err = producer.SendMessage("notic", "Login", msg); err != nil {
		return r, err
	}

	return r, nil
}
