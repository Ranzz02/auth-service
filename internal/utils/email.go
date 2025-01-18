package utils

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/Ranzz02/auth-service/config"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

var mailer *EmailHandler

type EmailHandler struct {
	h      *hermes.Hermes
	dialer *gomail.Dialer
}

func NewEmailHandler() {
	h := &hermes.Hermes{
		Product: hermes.Product{
			Name:      "DiscGolf App",
			Link:      "https://rasmus-raiha.com",
			Logo:      "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fusdgc.com%2Fwp-content%2Fuploads%2F2015%2F10%2FLogo-web.jpg&f=1&nofb=1&ipt=d778b5d587c0f2ef66dfb94a00ea1339c0c45452d9a68fb2c3a4590d8478280c&ipo=images",
			Copyright: "Copyright © 2024 Disc Golf App. All rights reserved",
		},
	}

	// Create dialer
	config := config.NewEnvConfig()
	d := gomail.NewDialer(config.SmtpServer, config.SmtpPort, config.SmtpUser, config.SmtpPassword)

	mailer = &EmailHandler{
		h:      h,
		dialer: d,
	}

}

type smtpAuthentication struct {
	SenderIdentity string
	SenderEmail    string
}

type sendOptions struct {
	To      string
	Subject string
}

func (h *EmailHandler) send(smtpConfig smtpAuthentication, options sendOptions, htmlBody string, txtBody string) error {
	if smtpConfig.SenderIdentity == "" {
		return errors.New("SMTP sender identity is empty")
	}

	if smtpConfig.SenderEmail == "" {
		return errors.New("SMTP sender email is empty")
	}

	if options.To == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	return h.dialer.DialAndSend(m)
}

// Confirm account email
//
// Options struct
type ConfirmMailOptions struct {
	Username string
	To       string
	Link     string
}

// Function itself
func SendConfirmEmail(options ConfirmMailOptions) (bool, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Name: options.Username,
			Intros: []string{
				"Welcome to discgolf_app! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with discgolf_app, please click here:",
					Button: hermes.Button{
						Color: "#22bc66",
						Text:  "Confirm your account",
						Link:  options.Link,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help!",
			},
		},
	}

	emailBody, err := mailer.h.GenerateHTML(email)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	config := config.NewEnvConfig()

	auth := smtpAuthentication{
		SenderEmail:    config.SenderEmail,
		SenderIdentity: config.SenderIdentity,
	}

	conf := sendOptions{
		To:      options.To,
		Subject: "DISC GOLF APP",
	}

	if err := mailer.send(auth, conf, emailBody, ""); err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	fmt.Println("Email sent")
	return true, nil
}
