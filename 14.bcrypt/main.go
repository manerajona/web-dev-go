package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password)); err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

func main() {
	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	if err = comparePassword(pass, hashedPass); err != nil {
		log.Fatalln("Not logged in")
	}
	log.Println("Logged in!")
}
