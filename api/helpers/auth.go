package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword Hashes and salts a given string
func HashPassword(pwd string) (string, error) {

	passwordInBytes := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(passwordInBytes, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}
