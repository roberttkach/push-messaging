package main

import (
	"log"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(phone, message string) {
	accountSid := twilioAccountSid
	authToken := twilioAuthToken
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(twilioPhoneNum)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
	}
}
