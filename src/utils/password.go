package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plaintText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintText), 10)
	return string(bytes), err
}

func ComparePassword(plaintText, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintText))
	return err == nil, err
}
