package routes

import (
	"fmt"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	token := c.FormValue("token")
	content, err := utils.ReadToken(token)
	if err != nil {
		fmt.Println(err)
	}
	if token != content {
		return c.SendStatus(401)
	} else {
		return c.SendStatus(200)
	}
	return nil
}
