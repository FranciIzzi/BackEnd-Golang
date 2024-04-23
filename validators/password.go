package validators

import (
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordWithHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ValidatePassword(password, email string) (bool, string) {

	if len(password) < 8 {
		return false, "Password too short"
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, "Password doesn't contain at least one number"
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, "Password doesn't contain at least one maiusc letter"
	}

	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		return false, "Password doesn't contain at least one special character"
	}

	//  provide a list of basic and not secure password to avoid
	bannedPasswords := []string{"password", "qwerty", "12345678"}
	for _, b := range bannedPasswords {
		if strings.ToLower(password) == b {
			return false, "Password too easy, not secure"
		}
	}

	if strings.Contains(strings.ToLower(password), strings.ToLower(strings.Split(email, "@")[0])) {
		return false, "Password similar to the email, not secure"
	}

	return true,"Password secure !"
}
