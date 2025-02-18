package routes

import (
	handlers "go/starter-kit/api/src/handlers/user"
	"go/starter-kit/api/src/middleware"
	repositories "go/starter-kit/api/src/repositories/user"
	services "go/starter-kit/api/src/services/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRoutes(app *fiber.App, db *sqlx.DB, path string) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(*repo)
	handler := handlers.NewUserHandler(*service)

	group := app.Group(path, middleware.AuthMiddleWare)
	group.Get("/", handler.GetUserById)
}
