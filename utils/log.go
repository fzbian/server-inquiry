package utils

import (
	"bufio"
	"fmt"
	"github.com/fzbian/server-inquiry/enums"
	"os"
	"time"
)

/*
Log records one or more given commands and saves them to a .txt file
Parameters:
  - string array: takes the string or several strings and saves them

Return:
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
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
			fmt.Println(Problem(enums.CantCloseFile, err))
		}
	}(file)

	descriptor := bufio.NewWriter(file)

	for _, log := range logs {
		_, err := descriptor.WriteString("[" + FormatDate + "]: " + log + "\n")
		if err != nil {
			fmt.Println(Problem(enums.CantWriteFile, err))
		}
	}

	return descriptor.Flush()
}
