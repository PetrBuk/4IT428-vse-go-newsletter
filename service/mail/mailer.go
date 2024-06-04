package mail

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"

	svcmodel "vse-go-newsletter-api/service/model"
)

func sendMail(recepient string, subject string, content string) error {

	from := os.Getenv("EMAIL_ADDRESS")
	pw := os.Getenv("EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte("To: " + recepient + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "MIME-version: 1.0;\r\n" +
        "Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
        content)

	auth := smtp.PlainAuth("", from, pw, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recepient}, msg)
	if err != nil {
		return err
	}
	return nil

}

func SendConfirmationRequestMail(email string, newsletterId string) error {
	baseUrl := os.Getenv("BASE_URL")

	// Send confirmation email
	templateData := struct {
		ConfirmLink string
		UnsubscribeLink string
	}{
		ConfirmLink: baseUrl + "/api/v1/newsletters/" + newsletterId + "/confirm?email=" + email,
		UnsubscribeLink: baseUrl + "/api/v1/newsletters/" + newsletterId + "/unsubscribe?email=" + email,
	}

	t, err := template.ParseFiles("templates/confirm_request.html")

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		return err
	}

	html := buf.String()

	errEmail := sendMail(email, "Confirm subscription | VŠE Newsletter", html)

	if errEmail != nil {
		return errEmail
	}

	return nil
}

func SendConfirmationMail(email string, newsletterId string) error {
	baseUrl := os.Getenv("BASE_URL")

	// Send confirmation email
	templateData := struct {
		SubscriberEmail string
		UnsubscribeLink string
	}{
		SubscriberEmail: email,
		UnsubscribeLink: baseUrl + "/api/v1/newsletters/" + newsletterId + "/unsubscribe?email=" + email,
	}

	t, err := template.ParseFiles("templates/confirmed.html")

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		return err
	}

	html := buf.String()

	errEmail := sendMail(email, "Subscription confirmed | VŠE Newsletter", html)

	if errEmail != nil {
		return errEmail
	}

	return nil
}

func SendNewPostMail(subscribers []string, post *svcmodel.Post) error {
	baseUrl := os.Getenv("BASE_URL")

	for _, subscriber := range subscribers {
		// Send confirmation email
		templateData := struct {
			NewsletterContent string
			UnsubscribeLink  string
		}{
			NewsletterContent: post.Content,
			UnsubscribeLink:  baseUrl + "/api/v1/newsletters/" + post.NewsletterId.String() + "/unsubscribe?email=" + subscriber,
		}

		t, err := template.ParseFiles("templates/post.html")

		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		if err = t.Execute(buf, templateData); err != nil {
			return err
		}

		html := buf.String()

		errEmail := sendMail(subscriber, post.Title + " | VŠE Newsletter", html)

		if errEmail != nil {
			return errEmail
		}
	}

	return nil
}
