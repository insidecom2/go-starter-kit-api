package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare(ctx *fiber.Ctx) (err error) {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_TOKEN_KEY"))},
		TokenLookup: "header:Authorization", // Look for the token in headers
		AuthScheme:  "Bearer",               // Ensure it's a Bearer token
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or missing token",
			})
		}, SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			userData := claims["sub"].(map[string]interface{})
			ctx.Locals("email", userData["email"])
			ctx.Locals("user_id", userData["id"])

			return ctx.Next()
		},
	})(ctx)
}
