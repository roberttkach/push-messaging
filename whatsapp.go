package main

import (
	"log"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendWhatsApp(phone, message string) {
	accountSid := twilioAccountSid
	authToken := twilioAuthToken
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("whatsapp:" + phone)
	params.SetFrom(twilioWhatsAppNum)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending WhatsApp message: %v", err)
	}
}
