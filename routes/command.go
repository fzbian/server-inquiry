package routes

import (
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
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
			utils.ExecWindows(c, command)
			return nil
		} else if so == "linux" {
			utils.ExecLinux(c, command)
			return nil
		}
	}

	return nil
}
