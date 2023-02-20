package main

import (
	"github.com/fzbian/server-inquiry/routes"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName:               "Server Inquiry",
		DisableStartupMessage: true,
	})

	api := app.Group("/api")

	api.Get("/health", routes.Health)
	api.Get("/command", routes.Command)

	token := utils.GenerateToken()
	res, err := utils.SaveToken(token)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
