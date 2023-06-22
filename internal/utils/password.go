package utils

import (
	"fmt"
	"github.com/matthewhartstonge/argon2"
)

func HashPassword(password string) (string, error) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", fmt.Errorf("unable to hash password: %v", err.Error())
	}

	return string(encoded), err
}

func VerifyPassword(encoded string, password string) (bool, error) {
	return argon2.VerifyEncoded([]byte(password), []byte(encoded))
}
