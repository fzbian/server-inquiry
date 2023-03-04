package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"os/exec"
)

type Token struct {
	AccessToken string
}

var Tokens *Token

/*
GenerateToken using mathematical functions generates a code that will be used as a token.
Return:
  - string: The token generated in string format
*/
func GenerateToken() string {
	token := fmt.Sprintf("%s", uuid.NewV4())

	Tokens = &Token{token}

	return token
}

/*
VerifyToken uses uuid to generate the token in hash format
Return:
  - string: returns the token in string format
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func VerifyToken(token string) bool {
	return token == Tokens.AccessToken
}

func KillToken() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	Tokens = &Token{""}

	fmt.Println("I can't start the server, the port is in use.\nThis token was mistakenly removed from the server")
}
