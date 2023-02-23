package main

import (
	"github.com/fzbian/server-inquiry/routes"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:               "Server Inquiry",
		DisableStartupMessage: true,
	})

	api := app.Group("/api")

	api.Get("/health", routes.Health)
	api.Get("/command", routes.Command)

	utils.GenerateToken()
	err := app.Listen(":8080")
	if err.Error() == "failed to listen: listen tcp4 :8080: bind: address already in use" {
		utils.KillToken()
	}
}
