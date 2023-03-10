package routes

import (
	"fmt"
	"strings"

	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Error interface{}
}

/*
Command is in charge of obtaining the variables passed by the query and executing the necessary functions for its operation.
Return:
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func Command(c *fiber.Ctx) error {
	QueryToken := c.FormValue("token")
	QueryCommand := c.FormValue("cmd")

	if QueryCommand == "" || QueryToken == "" {
		return c.Status(400).JSON(Response{
			Error: enums.ParamsRequired,
		})
	}

	tokenExist := utils.VerifyToken(QueryToken)

	if !tokenExist {
		return c.Status(401).JSON(Response{
			Error: enums.WrongToken,
		})

	}

	for _, forbidden := range enums.Debian {
		if strings.Contains(QueryCommand, forbidden) {
			DangerousCommandMessage := fmt.Sprintf(enums.DangerousCommand, QueryCommand)
			return c.Status(405).SendString(DangerousCommandMessage)
		}
	}

	res, err := utils.Exec(c, QueryCommand)
	if err != nil && err.Error() == "exit status 127" {
		err := c.SendString(utils.Indicate(enums.CommandNotFound, QueryCommand))
		if err != nil {
			fmt.Println(utils.Indicate(enums.CantSendOutputCmd, err))
		}
	} else {
		err := c.SendString(res)
		if err != nil {
			fmt.Println(utils.Indicate(enums.CantSendOutputCmd, err))
		}
	}
	return nil
}
