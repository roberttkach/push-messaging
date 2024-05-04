package main

import (
	"fmt"
	"os"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbName     = os.Getenv("DB_NAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbSSLMode  = os.Getenv("DB_SSLMODE")
	dbHost     = os.Getenv("DB_HOST")

	emailFrom     = os.Getenv("EMAIL_FROM")
	emailPassword = os.Getenv("EMAIL_PASSWORD")
	smtpHost      = os.Getenv("SMTP_HOST")
	smtpPort      = os.Getenv("SMTP_PORT")

	twilioAccountSid  = os.Getenv("TWILIO_ACCOUNT_SID")
	twilioAuthToken   = os.Getenv("TWILIO_AUTH_TOKEN")
	twilioPhoneNum    = os.Getenv("TWILIO_PHONE_NUMBER")
	twilioWhatsAppNum = os.Getenv("TWILIO_WHATSAPP_NUMBER")

	telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")

	viberAPIKey = os.Getenv("VIBER_API_KEY")
)

type Recipient struct {
	Phone string
	Email string
}

func GetDBConnectionString() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s", dbUser, dbName, dbPassword, dbSSLMode, dbHost)
}
