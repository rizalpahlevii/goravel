package helpers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
