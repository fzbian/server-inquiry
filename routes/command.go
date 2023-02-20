package routes

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

/*
Command is in charge of obtaining the variables passed by the query and executing the necessary functions for its operation.
Return:
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func Command(c *fiber.Ctx) error {
	token := c.FormValue("token")
	content, err := utils.ReadToken()
	if err != nil {
		return err
	}
	if token != content {
		err := c.SendStatus(401)
		if err != nil {
			return err
		}
	} else {
		command := c.FormValue("cmd")
		so := c.FormValue("so")

		if so == "windows" {
			for _, forbidden := range enums.PowerShell {
				if strings.Contains(command, forbidden) {
					DangerousCommandMessage := fmt.Sprintf("The %s command can be dangerous to use remotely.", command)
					return c.Status(405).SendString(DangerousCommandMessage)
				}
			}

			err := utils.ExecWindows(c, command)
			if err != nil {
				return err
			}
			return nil
		} else if so == "linux" {
			for _, forbidden := range enums.Debian {
				if strings.Contains(command, forbidden) {
					DangerousCommandMessage := fmt.Sprintf("The %s command can be dangerous to use remotely.", command)
					return c.Status(405).SendString(DangerousCommandMessage)
				}
			}

			err := utils.ExecLinux(c, command)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
