package smtp

import (
	"bytes"
	"github.com/haski007/photo-landing/internal/app/web-server/repository"
	"html/template"
	"net/smtp"
)

type SMTPRepository struct {
	User string
	Pass string
	Host string
	Port string
}

func NewSMTPRepository(user, pass, host, port string) *SMTPRepository {
	return &SMTPRepository{
		User: user,
		Pass: pass,
		Host: host,
		Port: port,
	}
}

func (r *SMTPRepository) SendMail(recipient string, data *repository.EmailData) error {
	tmpl, err := template.ParseFiles("templates/email.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(mimeHeaders))

	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", r.User, r.Pass, r.Host)

	err = smtp.SendMail(r.Host+":"+r.Port, auth, r.User, []string{}, body.Bytes())

	return err
}
