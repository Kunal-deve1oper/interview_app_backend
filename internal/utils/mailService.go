package utils

import (
	"crypto/tls"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(subject string, receiver string, message string) {
	go func() {
		m := gomail.NewMessage()

		m.SetHeader("From", os.Getenv("MAIL_SERVICE_EMAIL"))
		m.SetHeader("To", receiver)

		m.SetHeader("Subject", subject)
		m.SetBody("text/html", message)

		d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("MAIL_SERVICE_EMAIL"), os.Getenv("MAIL_SERVICE_PASSWORD"))

		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := d.DialAndSend(m); err != nil {
			log.Print(err)
		}
	}()
}
