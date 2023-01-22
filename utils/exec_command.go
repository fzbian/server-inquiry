package utils

import (
	"github.com/gofiber/fiber/v2"
	"os/exec"
	"sync"
)

var (
	wg sync.WaitGroup
)

func ExecWindows(c *fiber.Ctx, command string) error {
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

func ExecLinux(c *fiber.Ctx, command string) error {
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
