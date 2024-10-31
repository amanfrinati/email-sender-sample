package main

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/mail.v2"
)

var (
	from     = "email"
	password = "password"
	to       = "alberto.manfrinati@uqido.com"
	hostname = "smtp" // Your SMTP server address
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Gomail test subject")
	m.SetBody("text/plain", "This is Gomail test body")

	d := gomail.NewDialer(hostname, 465, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: hostname}

	// Send N emails in a loop
	for i := 0; i < 20; i++ {
		log.Printf("Sending email...")

		if err := d.DialAndSend(m); err != nil {
			log.Printf("%s", err)
			panic(err)
		} else {
			log.Printf("Email sent!")
		}
	}
}
