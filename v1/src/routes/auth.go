package routes

import (
	"go/starter-kit/api/src/configs"
	handlers "go/starter-kit/api/src/handlers/auth"
	repositories "go/starter-kit/api/src/repositories/auth"
	services "go/starter-kit/api/src/services/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func AuthRoutes(app *fiber.App, db *sqlx.DB, config *configs.Config, path string) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(*authRepo, config)
	authHandlers := handlers.NewAuthHandler(*authService)

	group := app.Group(path)
	group.Post("/register", authHandlers.Register)
	group.Post("/login", authHandlers.Login)

}
