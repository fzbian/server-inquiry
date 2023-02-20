package utils

import (
	"github.com/gofiber/fiber/v2"
	"os/exec"
	"sync"
)

var (
	wg sync.WaitGroup
)

/*
ExecLinux run a command in the linux operating system
Parameters:
  - string (command): takes the command dictated by the request

Return:
  - string: the answer given by the system
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func ExecLinux(c *fiber.Ctx, command string) error {
	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	err = c.SendString(string(out))
	if err != nil {
		return err
	}
	return nil
}

/*
ExecWindows run a command in the windows operating system
Parameters:
  - string (command): takes the command dictated by the request

Return:
  - string: the answer given by the system
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func ExecWindows(c *fiber.Ctx, command string) error {
	cmd := exec.Command("powershell", "-Command", command)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	err = c.SendString(string(out))
	if err != nil {
		return err
	}
	return nil
}
