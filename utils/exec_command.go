package utils

import (
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/gofiber/fiber/v2"
	"os/exec"
	"sync"
)

var (
	wg sync.WaitGroup
)

/*
Exec run a command in the linux operating system
Parameters:
  - string (command): takes the command dictated by the request

Return:
  - string: the answer given by the system
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func Exec(c *fiber.Ctx, command string) error {
	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(Problem(enums.CantReadOutputCmd, err))
	}
	err = c.SendString(string(out))
	if err != nil {
		fmt.Println(Problem(enums.CantSendOutputCmd, err))
	}
	return nil
}
