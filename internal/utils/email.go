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
	h *hermes.Hermes
}

func NewEmailHandler() {
	h := &hermes.Hermes{
		Product: hermes.Product{
			Name: "DiscGolf App",
			Link: "https://rasmus-raiha.com",
			Logo: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fusdgc.com%2Fwp-content%2Fuploads%2F2015%2F10%2FLogo-web.jpg&f=1&nofb=1&ipt=d778b5d587c0f2ef66dfb94a00ea1339c0c45452d9a68fb2c3a4590d8478280c&ipo=images",
		},
	}

	mailer = &EmailHandler{
		h: h,
	}
}

type smtpAuthentication struct {
	Server         string
	Port           int
	SMTPUser       string
	SMTPPassword   string
	SenderIdentity string
	SenderEmail    string
}

type sendOptions struct {
	To      string
	Subject string
}

func (h *EmailHandler) send(smtpConfig smtpAuthentication, options sendOptions, htmlBody string, txtBody string) error {
	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}
	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

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

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)

	return d.DialAndSend(m)
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
		Server:         config.SmtpServer,
		Port:           config.SmtpPort,
		SenderEmail:    config.SenderEmail,
		SenderIdentity: config.SenderIdentity,
		SMTPUser:       config.SmtpUser,
		SMTPPassword:   config.SmtpPassword,
	}

	conf := sendOptions{
		To:      options.To,
		Subject: "DISC GOLF APP",
	}

	if err := mailer.send(auth, conf, emailBody, ""); err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	fmt.Println("Sent email")
	return true, nil
}
