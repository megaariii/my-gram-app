package helper

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(p string) (string, error) {
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	return string(hash), err
}

func ComparePass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}