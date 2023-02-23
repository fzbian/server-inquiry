package routes

import (
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

	tokenExist := utils.VerifyToken(token)
	if !tokenExist {
		return c.SendStatus(401)
	}

	return c.SendStatus(200)
}
