package routes

import (
	handlers "go/starter-kit/api/src/handlers/auth"
	repositories "go/starter-kit/api/src/repositories/auth"
	services "go/starter-kit/api/src/services/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func AuthRoutes(app *fiber.App, db *sqlx.DB, path string) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(*authRepo)
	authHandlers := handlers.NewAuthHandler(*authService)

	group := app.Group(path)
	group.Post("/register", authHandlers.Register)

}
