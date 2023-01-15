package routes

import (
	"github.com/fzbian/server-inquiry/utils"
	"github.com/gofiber/fiber/v2"
	"os/exec"
	"sync"
)

func Command(c *fiber.Ctx) error {
	token := c.FormValue("token")
	content, err := utils.ReadToken(token)
	if err != nil {
		return err
	}
	if token != content {
		err := c.SendStatus(401)
		if err != nil {
			return err
		}
	} else {
		command := c.FormValue("cmd")
		so := c.FormValue("so")
		var wg sync.WaitGroup

		if so == "windows" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				cmd := exec.Command("powershell", command)
				out, _ := cmd.Output()
				err := c.SendString(string(out))
				if err != nil {
					return
				}
			}()
			wg.Wait()
			return nil
		} else if so == "linux" {
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
	}

	return nil
}
