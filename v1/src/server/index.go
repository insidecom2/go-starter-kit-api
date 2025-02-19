package server

import (
	"fmt"
	"go/starter-kit/api/src/configs"
	"go/starter-kit/api/src/routes"
	"go/starter-kit/api/src/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type SeverStruct struct {
	configs  *configs.Config
	database *sqlx.DB
}

func Init() *fiber.App {
	config := configs.Config{}
	server := &SeverStruct{
		configs:  &config,
		database: &sqlx.DB{},
	}
	server.loadEnv()
	server.ConnectDB()
	defer server.database.Close()
	app, _ := server.Routes()
	fmt.Println(">> Start API V1 <<")
	return app
}

func (s *SeverStruct) loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic("Cannot get .env file")
	}
	s.configs.DBConfig.Type = os.Getenv("DB_TYPE")
	s.configs.DBConfig.Host = os.Getenv("DB_HOST")
	s.configs.DBConfig.Port = os.Getenv("DB_PORT")
	s.configs.DBConfig.DBName = os.Getenv("DB_DATABASE")
	s.configs.DBConfig.User = os.Getenv("DB_USERNAME")
	s.configs.DBConfig.Password = os.Getenv("DB_PASSWORD")

	s.configs.SeverConfig.Port = os.Getenv("PORT")

	s.configs.JwtToken.Token = os.Getenv("JWT_TOKEN_KEY")
	s.configs.JwtToken.ReFreshToken = os.Getenv("JWT_REFRESH_TOKEN_KEY")

}

func (s *SeverStruct) ConnectDB() {

	postgresUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", s.configs.DBConfig.Type, s.configs.DBConfig.User, s.configs.DBConfig.Password, s.configs.DBConfig.Host, s.configs.DBConfig.Port, s.configs.DBConfig.DBName)
	db, err := sqlx.Connect("postgres", postgresUrl)

	if err != nil {
		log.Fatal(err)
	}
	s.database = db
	fmt.Println(">> Connected postgres database. <<")

}

func (s *SeverStruct) Routes() (app *fiber.App, err error) {
	app = fiber.New()
	app.Use(cors.New())

	prefix := "/v1"
	v1 := app.Group(prefix)
	v1.Get("/healthy", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(utils.ResponseStruct[interface{}]{
			Message: "healthy",
		})
	})
	// route auth
	routes.AuthRoutes(app, s.database, s.configs, prefix+"/auth")
	routes.UserRoutes(app, s.database, prefix+"/user")

	err = app.Listen(":" + s.configs.SeverConfig.Port)
	if err != nil {
		log.Fatal(err)
	}
	return app, nil
}
