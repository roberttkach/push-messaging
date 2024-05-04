package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendTelegram(phone, message string) {
	botToken := telegramBotToken
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Printf("Error creating Telegram bot: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(0, message)
	msg.ParseMode = "HTML"

	chatID, err := getChatIDByPhone(phone)
	if err != nil {
		log.Printf("Error getting chat ID by phone: %v", err)
		return
	}
	msg.ChatID = chatID

	_, err = bot.Send(msg)
	if err != nil {
		log.Printf("Error sending Telegram message: %v", err)
	}
}

func getChatIDByPhone(phone string) (int64, error) {
	botToken := telegramBotToken
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return 0, err
	}

	user, err := bot.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{
		UserID: phone,
		Limit:  1,
	})
	if err != nil {
		return 0, err
	}

	if user.TotalCount == 0 {
		return 0, fmt.Errorf("user not found for phone number: %s", phone)
	}

	return user.Photos[0][0].FileID, nil
}
