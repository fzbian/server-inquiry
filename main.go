package main

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/routes"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/exec"
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
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("cls error: %s\n", err)
		}
		fmt.Println("The token was deleted because the server was not started correctly.")
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
