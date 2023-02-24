package utils

import "fmt"

func Problem(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return fmt.Sprintf("(server inquiry) [error]: %s", message)
}

func Indicate(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s", message)
}
