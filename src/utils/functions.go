package utils

import "strings"

func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".com")
}

func HasValidSize(str string, validSize int) bool {
	return len(str) >= validSize
}
