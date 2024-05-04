package main

import (
	"log"

	"github.com/viber/go-viber"
)

func SendViber(phone, message string) {
	apiKey := viberAPIKey
	sender := viber.Sender{
		Name: "Your Viber Bot",
	}

	v := viber.New(apiKey, sender, nil)

	userID, err := getUserIDByPhone(phone)
	if err != nil {
		log.Printf("Error getting Viber user ID by phone: %v", err)
		return
	}

	msg := viber.TextMessage{
		Receiver: userID,
		Text:     message,
	}

	_, err = v.SendMessage(msg)
	if err != nil {
		log.Printf("Error sending Viber message: %v", err)
	}
}

func getUserIDByPhone(phone string) (string, error) {
	apiKey := viberAPIKey
	sender := viber.Sender{
		Name: "Your Viber Bot",
	}

	v := viber.New(apiKey, sender, nil)

	user, err := v.GetUserDetails(phone)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
