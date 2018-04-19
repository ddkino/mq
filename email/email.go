package email

import (
	"net/smtp"

	"strconv"
)

type Email struct {
	from    string
	to      []string
	subject string
	body    string
}

func New() *Email {
	return new(Email)
}

func (e *Email) Init(from string, to []string, subject string, body string) {
	e.from = from
	e.to = to
	e.subject = subject
	e.body = body
}

func (e *Email) SendEmail(ENV string) error {
	/**
	------------- SENDMAIL --------------
	*/
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "subject: " + e.subject + "\n"
	msg := []byte(subject + mime + "\n" + e.body)

	smtpServer := NewSmtpConfig(ENV)
	addr := smtpServer.Host + ":" + strconv.Itoa(smtpServer.Port)
	auth := smtp.PlainAuth("", smtpServer.Username, smtpServer.Password, smtpServer.Host)

	if err := smtp.SendMail(addr, auth, e.from, e.to, msg); err != nil {
		return err
	}
	return nil
}
