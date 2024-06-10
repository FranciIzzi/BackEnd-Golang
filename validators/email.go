package validators

import (
	"log"
	"net/smtp"
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

func SendWelcomeEmail(email, token string) error {
	// passare queste variabili nel .env
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := "franci99.izzi@gmail.com"
	smtpPass := "eall yort pqxu yhkw"

	msg := []byte("To: " + email + "\r\n" +
		"Subject: Benvenuto!\r\n" +
		"\r\n" +
		"Ciao! Per impostare la tua password, visita il seguente link:\r\n" +
		"http://localhost:8000/impact/v1/api/rest/pw/reset-token=" + token + "\r\n")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{email}, msg)
	if err != nil {
		log.Fatal("Errore nell'invio dell'email:", err)
		return err
	}
	return nil
}
