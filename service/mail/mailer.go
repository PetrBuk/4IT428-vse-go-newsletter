package mail

import (
	"net/smtp"
	"os"
)

func SendMail(subscribers []string, content string) error {

	from := os.Getenv("EMAIL_ADDRESS")
	pw := os.Getenv("EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(content)

	auth := smtp.PlainAuth("", from, pw, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, subscribers, message)
	if err != nil {
		return err
	}
	return nil

}
