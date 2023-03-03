package main

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/routes"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
)

var PORT = "3000"

func main() {
	app := fiber.New(fiber.Config{
		AppName:               "Server Inquiry",
		DisableStartupMessage: true,
	})

	api := app.Group("/api")

	api.Get("/health", routes.Health)
	api.Get("/command", routes.Command)

	token := utils.GenerateToken()
	fmt.Printf("Token: %s", token)
	err := app.Listen(":" + PORT)
	if err.Error() == utils.Indicate(enums.PortAlreadyUsed, PORT) {
		utils.KillToken()
	}
}
