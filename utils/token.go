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
SaveToken saves the token to a local file for later reading
Parameters:
  - string (token): takes the generated token, the GenerateToken function is used to generate it.

Return:
  - string: returns a message informing the user what the token is, and is saved in a token.yml file.
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
/*func SaveToken(token string) (string, error) {
		tokenYAML, err := yaml.Marshal(content)
	if err != nil {
		return "", err
	}

	file, err := os.Create("token.yml")
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.Write(tokenYAML)
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Generated token, remember not to share it: %s", token)
	return res, nil
}*/

/*
VerifyToken reads the token.yml file to access the AccessToken variable and returns it
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
