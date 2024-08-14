package util

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

// IsValidEmail checks if the provided email address is valid using a regular expression.
func IsValidEmail(email string) bool {
	// Simple regex to check the email format
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// IsValidPassword checks if the password meets the minimum security criteria.
func IsValidPassword(password string) bool {
	// Password should be at least 8 characters long and contain at least one number and one letter.
	if len(password) < 8 {
		return false
	}

	var hasLetter, hasDigit bool
	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z':
			hasLetter = true
		case '0' <= char && char <= '9':
			hasDigit = true
		}
	}

	return hasLetter && hasDigit
}
