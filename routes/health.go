package routes

import (
	"fmt"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
)

/*
Health checks if the token provided in the "token" parameter is valid.
401 if the token is not valid
200 if token is valid, is authorized
Return:
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func Health(c *fiber.Ctx) error {
	token := c.FormValue("token")
	content, err := utils.ReadToken()
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
