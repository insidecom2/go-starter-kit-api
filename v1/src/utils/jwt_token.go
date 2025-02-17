package utils

import (
	"go/starter-kit/api/src/consts"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	Key string
}

func NewJwtToken(key string) *JwtToken {
	return &JwtToken{
		Key: key,
	}
}

func (j *JwtToken) HashJwtToken(body consts.UserResponse) (r string, err error) {
	secretKey := []byte(j.Key)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": body,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	r, err = token.SignedString(secretKey)
	if err != nil {
		return r, err
	}
	return r, nil
}
