package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"math/rand"
	"os"
	"time"
)

type Token struct {
	AccessToken string
}

/*
GenerateToken using mathematical functions generates a code that will be used as a token.
Return:
  - string: The token generated in string format
*/
func GenerateToken() string {
	var (
		ABC = "abcdefghijklmnñopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)

	rand.Seed(time.Now().UnixNano())
	var letters []rune
	letters = []rune(ABC)
	b := make([]rune, 50)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/*
SaveToken saves the token to a local file for later reading
Parameters:
  - string (token): takes the generated token, the GenerateToken function is used to generate it.

Return:
  - string: returns a message informing the user what the token is, and is saved in a token.yml file.
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func SaveToken(token string) (string, error) {
	content := &Token{AccessToken: token}

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
}

/*
ReadToken reads the token.yml file to access the AccessToken variable and returns it
Return:
  - string: returns the token in string format
  - error: if there is an error, it is returned so that when the function is used, it can be reported.
*/
func ReadToken() (string, error) {
	file, err := os.Open("token.yml")
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var tokenYAML []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		tokenYAML = append(tokenYAML, buf[:n]...)
	}

	var content Token
	err = yaml.Unmarshal(tokenYAML, &content)
	if err != nil {
		return "", err
	}

	return content.AccessToken, nil
}
