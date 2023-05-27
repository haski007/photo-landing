package handler

import (
	"log"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

// ProcessForm handles the submission of the form
func ProcessForm(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	// Send an email
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	to := []string{os.Getenv("RECIPIENT_EMAIL")}
	msg := []byte("To: " + os.Getenv("RECIPIENT_EMAIL") + "\r\n" +
		"Subject: New Message from " + name + "\r\n" +
		"\r\n" +
		"Name: " + name + "\r\nEmail: " + email + "\r\nMessage: " + message + "\r\n")

	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":587", auth, os.Getenv("SMTP_USERNAME"), to, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Render a success page (optional)
	c.HTML(200, "success.html", nil)
}
