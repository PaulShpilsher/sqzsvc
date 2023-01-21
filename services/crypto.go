package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		log.Println("bcrypt: unexpected error", err)
	}
	return err
}
