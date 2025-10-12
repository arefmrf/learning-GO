package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(pass1 string, pass2 string) bool {
	fmt.Println("====", pass1, pass2)
	err := bcrypt.CompareHashAndPassword([]byte(pass1), []byte(pass2))
	return err == nil
}
