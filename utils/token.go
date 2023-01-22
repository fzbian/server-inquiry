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

func SaveToken() (string, error) {
	token := GenerateToken()
	content := &Token{AccessToken: token}

	tokenYAML, err := yaml.Marshal(content)
	if err != nil {
		return "", err
	}

	file, err := os.Create("token.yml")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(tokenYAML)
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Generated token, remember not to share it: %s", token)
	return res, nil
}

func ReadToken() (string, error) {
	file, err := os.Open("token.yml")
	if err != nil {
		return "", err
	}
	defer file.Close()

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
