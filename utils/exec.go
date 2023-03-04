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

	//Debian
	cmd := exec.Command("/bin/bash", "-c", command)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(Problem(enums.CantReadOutputCmd, err))
	}

	// Saves logs
	err = Log(command)
	if err != nil {
		fmt.Println("error 1: ", err.Error())
	}

	// Returns the output in string format
	err = c.SendString(string(out))
	if err != nil {
		fmt.Println(Problem(enums.CantSendOutputCmd, err))
	}
	return nil
}
