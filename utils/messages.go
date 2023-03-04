package utils

import (
	"fmt"
	"io"
	"net/http"
)

func Problem(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return fmt.Sprintf("(server inquiry) [error]: %s", message)
}

func Indicate(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s", message)
}

func GetIP() string {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := io.ReadAll(res.Body)
	return string(ip)
}
