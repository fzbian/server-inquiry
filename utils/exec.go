package utils

import (
	"bufio"
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/exec"
	"sync"
	"time"
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
	// PowerShell
	cmd := exec.Command("powershell", "-Command", command)
	// Debian
	//cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(Problem(enums.CantReadOutputCmd, err))
	}

	// Saves logs
	err = Log(command)
	if err != nil {
		fmt.Println("error 1: ", err.Error())
	}

	err = c.SendString(string(out))
	if err != nil {
		fmt.Println(Problem(enums.CantSendOutputCmd, err))
	}
	return nil
}

func Log(logs ...string) error {
	Date := time.Now()
	FormatDate := Date.Format("2006-01-02 15:04:05")

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error: ", err)
		}
	}(file)

	descriptor := bufio.NewWriter(file)

	for _, log := range logs {
		_, err := descriptor.WriteString("[" + FormatDate + "]: " + log + "\n")
		if err != nil {
			return err
		}
	}

	return descriptor.Flush()
}
