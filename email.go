package main

import (
	"log"
	"net/smtp"
)

func SendEmail(email, message string) {
	from := emailFrom
	password := emailPassword
	to := []string{email}
	smtpHost := smtpHost
	smtpPort := smtpPort

	msg := []byte("Subject: Message\r\n\r\n" + message)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
	}
}
