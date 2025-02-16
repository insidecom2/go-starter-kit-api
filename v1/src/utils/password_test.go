package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswordWithString(t *testing.T) {
	inputs := []struct {
		value        string
		hashedLength int
		wantError    bool
	}{
		{
			value:        "1234",
			hashedLength: 60,
			wantError:    false,
		},
		{
			value:        "",
			hashedLength: 60,
			wantError:    false,
		},
	}
	for _, input := range inputs {
		if !input.wantError {
			hashPass, _ := HashPassword(input.value)
			assert.Equal(t, len(hashPass), input.hashedLength, "")
		}
	}
}

func TestComparePassword(t *testing.T) {
	mocks := []struct {
		hashPassword string
		password     string
		expectValue  bool
	}{
		{
			hashPassword: "$2a$10$dkKIyYQTnxX8/qLJpIZLlu7r7R20.QLGmHLu6PxWKAIbDkxgQO.6e",
			password:     "1234",
			expectValue:  true,
		},
		{
			hashPassword: "$2a$10$n5jElozRxj4pvmIKFGhUaOIxbDbR9F35UJyKEdCY.gNkSu6JWtHtW",
			password:     "",
			expectValue:  true,
		},
		{
			hashPassword: "$2a$10$n5jElozRxj4pvmIKFGhUaOIxbDbR9F35UJyKEdCY.gNkSu6JWtHtW",
			password:     "2222",
			expectValue:  false,
		},
	}

	for _, mock := range mocks {
		ok := ComparePassword(mock.hashPassword, mock.password)
		assert.Equal(t, mock.expectValue, ok, "")
	}
}
