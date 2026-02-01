package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail checks if an email address is valid
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(strings.TrimSpace(email))
}

// ValidatePassword checks if a password meets requirements
// Requirements: at least 8 characters, 1 uppercase, 1 lowercase, 1 number
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}

	var hasUpper, hasLower, hasNumber bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUpper {
		return false, "Password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "Password must contain at least one lowercase letter"
	}
	if !hasNumber {
		return false, "Password must contain at least one number"
	}

	return true, ""
}

// SanitizeString trims and cleans a string
func SanitizeString(s string) string {
	return strings.TrimSpace(s)
}

// ValidateName checks if a name is valid (not empty, reasonable length)
func ValidateName(name string) bool {
	name = strings.TrimSpace(name)
	return len(name) >= 2 && len(name) <= 100
}

// ValidatePhone checks if a phone number is valid (basic check)
func ValidatePhone(phone string) bool {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return true // Phone is optional
	}
	// Remove common formatting characters
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '+' {
			return r
		}
		return -1
	}, phone)
	return len(cleaned) >= 10 && len(cleaned) <= 15
}

