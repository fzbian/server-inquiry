package utils

import (
	"fmt"
	"log"
	"os/exec"

	uuid "github.com/satori/go.uuid"
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
func GenerateToken() {
	token := fmt.Sprintf("%s", uuid.NewV4())

	Tokens = &Token{token}

	fmt.Println(token)
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
	x, err := exec.Command("clear").CombinedOutput()
	fmt.Print(string(x))
	if err != nil {
		log.Panic(err.Error())
	}

	Tokens = &Token{""}

	fmt.Println("This token was mistakenly removed from the server")
}
