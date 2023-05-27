package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(password string) bool {
	regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*_=+-]).{8,}$`)
	return regex.MatchString(password)
}

func HashPassword(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(passwordHash)
}

func ComparePassword(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
