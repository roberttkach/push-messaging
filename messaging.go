package main

func sendMessage(recipient Recipient, message string) {
	sendEmail(recipient.Email, message)
	sendSMS(recipient.Phone, message)
	sendWhatsApp(recipient.Phone, message)
	sendTelegram(recipient.Phone, message)
	sendViber(recipient.Phone, message)
}
