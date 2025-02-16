package main

import (
	"bytes"
	"encoding/json"
	"go/starter-kit/api/src/consts"
	"go/starter-kit/api/src/server"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAuthRoute(t *testing.T) {
	initApp := server.Init()

	tests := []struct {
		description  string
		requestBody  consts.RegisterRequest
		expectStatus int
	}{
		{
			description: "Valid input",
			requestBody: consts.RegisterRequest{
				UserEntity: consts.UserEntity{
					Email:    "a@a.com",
					Name:     "tom",
					Password: "1234",
				},
				ConfirmPassword: "1234",
			},
			expectStatus: fiber.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := initApp.Test(req)

			assert.Equal(t, test.expectStatus, resp.Status, "Ok")
		})
	}

}
