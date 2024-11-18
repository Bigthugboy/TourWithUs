package configs

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("no input value")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate encrypted password: %v", err)
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
