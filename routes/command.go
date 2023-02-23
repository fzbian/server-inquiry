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
	QueryToken := c.FormValue("token")
	QueryCommand := c.FormValue("cmd")

	LocalToken, err := utils.ReadToken()
	if err != nil {
		fmt.Println(utils.Problem(enums.CantReadToken, err))
	}

	if QueryToken != LocalToken {
		err := c.SendStatus(401)
		if err != nil {
			fmt.Println(utils.Problem(enums.CantSendCode, err))
		}
	}

	for _, forbidden := range enums.Debian {
		if strings.Contains(QueryCommand, forbidden) {
			DangerousCommandMessage := fmt.Sprintf(enums.DangerousCommand, QueryCommand)
			return c.Status(405).SendString(DangerousCommandMessage)
		}
	}

	err = utils.Exec(c, QueryCommand)
	if err != nil {
		fmt.Println(utils.Problem(enums.CantExecCommand, err))
	}
	return nil
}
