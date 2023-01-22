package routes

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

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
			err := utils.ExecLinux(c, command)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
