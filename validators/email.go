package validators

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) (bool, string) {
	if email == "" {
		return false, "L'email non può essere vuota."
	}

	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !regex.MatchString(email) {
		return false, "Il formato dell'email non è valido."
	}

	allowedDomains := []string{"gmail.com", "dedpartners.com"}
	for _, domain := range allowedDomains {
		if strings.HasSuffix(email, domain) {
			return true, ""
		}
	}

	return false, "Il dominio dell'email non è consentito."
}

