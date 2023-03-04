package main

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/routes"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var PORT = 8000
var app *fiber.App

func main() {

	for i := 0; i < 50; i++ {
		app = fiber.New(fiber.Config{
			AppName:               "Server Inquiry",
			DisableStartupMessage: true,
		})

		api := app.Group("/api")
		api.Get("/health", routes.Health)
		api.Get("/command", routes.Command)

		err := utils.ClearTerminal()
		if err != nil {
			fmt.Println(utils.Problem(enums.CantClearTerminal, err))
		}
		token := utils.GenerateToken()
		fmt.Printf("Server PORT: %d\nToken: %s\n", PORT, token)

		StringPORT := strconv.Itoa(PORT)
		err = app.Listen(":" + StringPORT)
		if err == nil {
			fmt.Println(utils.Problem(enums.CantStartServer, err))
			break
		}
		PORT++

		err = utils.ClearTerminal()
		if err != nil {
			fmt.Println(utils.Problem(enums.CantClearTerminal, err))
		}
	}
}
