package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println("Hashed password len: ", len(hashedPassword))
	return string(hashedPassword), nil
}

func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		log.Println("bcrypt: unexpected error", err)
	}
	return err
}
