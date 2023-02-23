package main

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
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

	Start()
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(utils.Problem(enums.ServerNotStartup, err))
	}
}

func Start() {
	token := utils.GenerateToken()
	res, err := utils.SaveToken(token)
	if err != nil {
		fmt.Println(utils.Problem(enums.TokenNotSaved, err))
	}
	fmt.Println(res)
}
