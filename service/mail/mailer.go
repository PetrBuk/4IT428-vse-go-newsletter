package mail

import (
	"net/smtp"
)

func SendMail(subscribers []string, content string) error {

	from := "gonewsvse@gmail.com"
	pw := "sfan xtcd snwv fkoc"

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
