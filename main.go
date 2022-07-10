package main

import (
	"fmt"

	/* should preceed any other repository imports */
	_ "github.com/bratushkadan/My-Fancy-Project/internal/init"

	"github.com/bratushkadan/My-Fancy-Project/internal/middlewares"
	"github.com/bratushkadan/My-Fancy-Project/internal/routes"
	"github.com/bratushkadan/My-Fancy-Project/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := prepareServer()

	app.Listen(fmt.Sprintf("localhost:%s", util.DefaultEnv("APP_PORT", "8090")))
}

func prepareServer() *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth")

	auth.Get("/token", routes.IssueToken)

	app.Use(middlewares.Auth())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	return app
}
