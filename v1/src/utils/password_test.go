package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHasher struct{}

func (m MockHasher) GenerateFromPassword([]byte, int) ([]byte, error) {
	return nil, errors.New("mock error")
}
func TestHashPasswordWithString(t *testing.T) {
	hashPass, _ := HashPassword("1234")
	assert.NotEqual(t, "", hashPass, "OK")

}
func TestHashPasswordWithEmptyString(t *testing.T) {
	hashPass, _ := HashPassword("")
	assert.NotEqual(t, "", hashPass, "empty string")
}

// func TestHashPasswordReturnError(t *testing.T) {

// 	monkey.Patch(bcrypt.GenerateFromPassword, func([]byte, int) ([]byte, error) {
// 		return nil, errors.New("error")
// 	})
// 	defer monkey.UnpatchAll()

// 	_, err := HashPassword("1234")
// 	assert.Error(t, err)
// 	assert.Equal(t, "error", err.Error())
// }
