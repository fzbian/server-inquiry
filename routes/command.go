package routes

import (
	"fmt"
	"strings"

	"github.com/fzbian/server-inquiry/enums"
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message interface{}
	Data    interface{}
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
			Message: "Error",
			Data:    "Token and Command are required",
		})
	}

	tokenExist := utils.VerifyToken(QueryToken)

	if !tokenExist {
		return c.Status(401).JSON(Response{
			Message: "Error",
			Data:    "This token not exist",
		})

	}

	for _, forbidden := range enums.Debian {
		if strings.Contains(QueryCommand, forbidden) {
			DangerousCommandMessage := fmt.Sprintf(enums.DangerousCommand, QueryCommand)
			return c.Status(405).SendString(DangerousCommandMessage)
		}
	}

	err := utils.Exec(c, QueryCommand)
	if err != nil {
		fmt.Println(utils.Problem(enums.CantExecCommand, err))
	}
	return nil
}
