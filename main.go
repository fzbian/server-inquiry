package main

import (
	"github.com/fzbian/server-inquiry/routes"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	api := app.Group("/api")

	api.Get("/health", routes.Health)
	api.Get("/command", routes.Command)

	res, err := utils.SaveToken()
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	app.Listen(":3000")
}
